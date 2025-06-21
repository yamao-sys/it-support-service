package businessservices

import (
	businessapi "apps/api/business"
	businessvalidators "apps/business/validators"
	models "apps/models"
	"errors"
	"fmt"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/volatiletech/null/v8"
	"gorm.io/gorm"
)

const toProjectPerPage = 5

type ToProjectService interface {
	FetchLists(pageToken int, startDate string, endDate string, supporterID int) (toProjects []ToProjectFields, nextPageToken int)
	Fetch(ID int, supporterID int) (toProject ToProjectFields, error error)
	CreatePlan(projectID int, requestParams *businessapi.PlanStoreWithStepsInput, supporterID int) (plan models.Plan, validatorErrors error, error error)
	MappingPlanWithStepsValidationErrorStruct(err error) businessapi.PlanWithStepsValidationError
}

type toProjectService struct {
	db *gorm.DB
}

type ToProjectFields struct {
	ID int
	Title string
	Description string
	StartDate time.Time
	EndDate time.Time
	MinBudget int
	MaxBudget int
	ProposalStatus businessapi.ToProjectProposalStatus
}

func NewToProjectService(db *gorm.DB) ToProjectService {
	return &toProjectService{db}
}

func (tps *toProjectService) FetchLists(pageToken int, startDate string, endDate string, supporterID int) (toProjects []ToProjectFields, nextPageToken int) {
	query := tps.db.Model(&models.Project{})

	if pageToken > 0 {
		query = query.Where("projects.id >= ?", pageToken)
	}
	if startDate != "" {
		query = query.Where("projects.start_date >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("projects.end_date <= ?", endDate)
	}
	query = query.Joins(
		"LEFT JOIN plans ON plans.project_id = projects.id AND plans.supporter_id = ?",
		supporterID,
	)

	query = query.Select(strings.Join(tps.toProjectFields(), ","))
	// NOTE: nextPageTokenの検出のため、1ページの件数+1を取得
	query = query.Limit(toProjectPerPage + 1)
	query = query.Limit(toProjectPerPage + 1)

	query.Scan(&toProjects)

	// NOTE: nextPageTokenのprojectをsliceから切り出し
	if len(toProjects) == toProjectPerPage + 1 {
		nextPageToken := toProjects[len(toProjects)-1].ID
		return toProjects[:len(toProjects)-1], nextPageToken
	}
	
	return toProjects, 0
}

func (tps *toProjectService) Fetch(ID int, supporterID int) (toProject ToProjectFields, error error) {
	query := tps.db.Model(&models.Project{})
	query = query.Joins(
		"LEFT JOIN plans ON plans.project_id = projects.id AND plans.supporter_id = ?",
		supporterID,
	)
	query = query.Select(strings.Join(tps.toProjectFields(), ","))

	query.First(&toProject, ID)
	if toProject.ID == 0 {
		return toProject, errors.New("not found")
	}
	
	return toProject, nil
}

func (tps *toProjectService) CreatePlan(projectID int, requestParams *businessapi.PlanStoreWithStepsInput, supporterID int) (plan models.Plan, validatorErrors error, error error) {
	// NOTE: プロジェクトの存在確認
	var project models.Project
	if err := tps.db.First(&project, projectID).Error; err != nil {
		fmt.Println("Project not found:", err)
		return models.Plan{}, nil, errors.New("project not found")
	}

	// NOTE: バリデーションチェック
	validatorErrors = businessvalidators.ValidatePlanWithSteps(requestParams)
	if validatorErrors != nil {
		return models.Plan{}, validatorErrors, nil
	}

	// NOTE: トランザクション開始
	tx := tps.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// NOTE: Planの作成
	plan = models.Plan{}
	plan.SupporterID = supporterID
	plan.ProjectID = projectID
	plan.Title = requestParams.Title
	plan.Description = requestParams.Description
	if requestParams.StartDate != nil {
		plan.StartDate = null.Time{Time: requestParams.StartDate.Time, Valid: true}
	}
	if requestParams.EndDate != nil {
		plan.EndDate = null.Time{Time: requestParams.EndDate.Time, Valid: true}
	}
	plan.UnitPrice = requestParams.UnitPrice
	plan.Status = models.PlanStatusTempraryCreating

	if err := tx.Create(&plan).Error; err != nil {
		tx.Rollback()
		return models.Plan{}, nil, err
	}

	// NOTE: PlanStepsの作成
	if requestParams.PlanSteps != nil && len(*requestParams.PlanSteps) > 0 {
		planSteps := make([]models.PlanStep, 0, len(*requestParams.PlanSteps))
		for _, stepInput := range *requestParams.PlanSteps {
			planStep := models.PlanStep{
				PlanID:      plan.ID,
				Title:       stepInput.Title,
				Description: stepInput.Description,
				Duration:    stepInput.Duration,
			}
			planSteps = append(planSteps, planStep)
		}

		if err := tx.Create(&planSteps).Error; err != nil {
			tx.Rollback()
			return models.Plan{}, nil, err
		}
	}

	// NOTE: トランザクションコミット
	if err := tx.Commit().Error; err != nil {
		return models.Plan{}, nil, err
	}

	// NOTE: 作成されたPlanにPlanStepsを含めて返す
	var finalPlan models.Plan
	if err := tps.db.Preload("PlanSteps").First(&finalPlan, plan.ID).Error; err != nil {
		return models.Plan{}, nil, err
	}

	return finalPlan, nil, nil
}

func (tps *toProjectService) MappingPlanWithStepsValidationErrorStruct(err error) businessapi.PlanWithStepsValidationError {
	var validationError businessapi.PlanWithStepsValidationError
	if err == nil {
		return validationError
	}

	if errors, ok := err.(validation.Errors); ok {
		// NOTE: レスポンス用の構造体にマッピング
		for field, err := range errors {
			messages := []string{err.Error()}
			switch field {
			case "title":
				validationError.Title = &messages
			case "description":
				validationError.Description = &messages
			case "startDate":
				validationError.StartDate = &messages
			case "endDate":
				validationError.EndDate = &messages
			case "unitPrice":
				validationError.UnitPrice = &messages
			case "planSteps":
				stepErrors := tps.parsePlanStepsValidationError(err)
				validationError.PlanSteps = &stepErrors
			}
		}
	}

	return validationError
}

func (tps *toProjectService) parsePlanStepsValidationError(err error) []businessapi.PlanStepValidationError {
	var stepErrors []businessapi.PlanStepValidationError
	
	// NOTE: PlanStepsのバリデーションエラーがvalidation.Errorsの場合、フィールド名で分類
	if validationErrors, ok := err.(validation.Errors); ok {
		stepError := businessapi.PlanStepValidationError{}
		
		for field, fieldErr := range validationErrors {
			messages := []string{fieldErr.Error()}
			
			switch field {
			case "title":
				stepError.Title = &messages
			case "description":
				stepError.Description = &messages
			case "duration":
				stepError.Duration = &messages
			}
		}
		
		stepErrors = append(stepErrors, stepError)
	} else {
		// NOTE: 単一エラーメッセージの場合（例：ステップ数が0など）
		stepError := businessapi.PlanStepValidationError{}
		messages := []string{err.Error()}
		stepError.Title = &messages
		stepErrors = append(stepErrors, stepError)
	}
	
	return stepErrors
}

func (tps *toProjectService) toProjectFields() []string {
	return []string{
		"projects.id",
		"projects.title",
		"projects.description",
		"projects.start_date",
		"projects.end_date",
		"projects.min_budget",
		"projects.max_budget",
		`CASE
		    WHEN plans.status IS NULL THEN 'NOT PROPOSED'
            WHEN plans.status = 0 THEN 'TEMPORARY CREATING'
            ELSE 'PROPOSED'
		 END AS proposal_status`,
	}
}
