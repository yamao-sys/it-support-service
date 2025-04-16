package businessservices

import (
	businessapi "apps/api/business"
	businessvalidators "apps/business/validators"
	models "apps/models/generated"
	"context"
	"database/sql"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const perPage = 5

type ProjectService interface {
	FetchLists(ctx context.Context, companyID int, pageToken int) (projects models.ProjectSlice, nextPageToken int, error error)
	Create(ctx context.Context, requestParams *businessapi.ProjectStoreInput, companyID int) (project models.Project, validatorErrors error, error error)
	Fetch(ctx context.Context, ID int) (project models.Project, error error)
	Update(ctx context.Context, requestParams *businessapi.ProjectStoreInput, ID int) (project models.Project, validatorErrors error, error error)
	MappingValidationErrorStruct(err error) businessapi.ProjectValidationError
}

type projectService struct {
	db *sql.DB
}

func NewProjectService(db *sql.DB) ProjectService {
	return &projectService{db}
}

func (ps *projectService) FetchLists(ctx context.Context, companyID int, pageToken int) (projects models.ProjectSlice, nextPageToken int, error error) {
	var conditions []qm.QueryMod
	conditions = append(conditions, qm.Where("company_id = ?", companyID))
	if pageToken > 0 {
		conditions = append(conditions, qm.Where("id >= ?", pageToken))
	}
	// NOTE: nextPageTokenの検出のため、1ページの件数+1を取得
	conditions = append(conditions, qm.Limit(perPage + 1))

	fetchedProjects, err := models.Projects(conditions...).All(ctx, ps.db)
	if err != nil {
		return models.ProjectSlice{}, 0, err
	}

	// NOTE: nextPageTokenのprojectをsliceから切り出し
	if len(fetchedProjects) == perPage + 1 {
		nextPageToken := fetchedProjects[len(fetchedProjects)-1].ID
		return fetchedProjects[:len(fetchedProjects)-1], nextPageToken, nil
	}
	
	return fetchedProjects, 0, nil
}

func (ps *projectService) Create(ctx context.Context, requestParams *businessapi.ProjectStoreInput, companyID int) (project models.Project, validatorErrors error, error error) {
	// NOTE: バリデーションチェック
	validatorErrors = businessvalidators.ValidateProject(requestParams)
	if validatorErrors != nil {
		return models.Project{}, validatorErrors, nil
	}

	project = models.Project{}
	project.CompanyID = companyID
	project.Title = *requestParams.Title
	project.Description = *requestParams.Description
	project.StartDate = requestParams.StartDate.Time
	project.EndDate = requestParams.EndDate.Time
	if requestParams.MinBudget != nil {
		project.MinBudget = null.Int{Int: *requestParams.MinBudget, Valid: true}
	}
	if requestParams.MaxBudget != nil {
		project.MaxBudget = null.Int{Int: *requestParams.MaxBudget, Valid: true}
	}
	project.IsActive = *requestParams.IsActive

	createErr := project.Insert(ctx, ps.db, boil.Infer())
	if createErr != nil {
		return models.Project{}, nil, createErr
	}

	return project, nil, nil
}

func (ps *projectService) Fetch(ctx context.Context, ID int) (project models.Project, error error) {
	fetchedProject, _ := models.Projects(qm.Where("id = ?", ID)).One(ctx, ps.db)
	if fetchedProject == nil {
		return models.Project{}, errors.New("not found")
	}
	
	return *fetchedProject, nil
}

func (ps *projectService) Update(ctx context.Context, requestParams *businessapi.ProjectStoreInput, ID int) (project models.Project, validatorErrors error, error error) {
	// NOTE: バリデーションチェック
	validatorErrors = businessvalidators.ValidateProject(requestParams)
	if validatorErrors != nil {
		return models.Project{}, validatorErrors, nil
	}

	doUpdateProject, _ := models.Projects(qm.Where("id = ?", ID)).One(ctx, ps.db)
	if doUpdateProject == nil {
		return models.Project{}, nil, errors.New("not found")
	}
	doUpdateProject.Title = *requestParams.Title
	doUpdateProject.Description = *requestParams.Description
	doUpdateProject.StartDate = requestParams.StartDate.Time
	doUpdateProject.EndDate = requestParams.EndDate.Time
	if requestParams.MinBudget != nil {
		doUpdateProject.MinBudget = null.Int{Int: *requestParams.MinBudget, Valid: true}
	}
	if requestParams.MaxBudget != nil {
		doUpdateProject.MaxBudget = null.Int{Int: *requestParams.MaxBudget, Valid: true}
	}
	doUpdateProject.IsActive = *requestParams.IsActive

	_, updateErr := doUpdateProject.Update(ctx, ps.db, boil.Infer())
	if updateErr != nil {
		return models.Project{}, nil, updateErr
	}
	project = *doUpdateProject

	return project, nil, nil
}

func (ps *projectService) MappingValidationErrorStruct(err error) businessapi.ProjectValidationError {
	var validationError businessapi.ProjectValidationError
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
			case "minBudget":
				validationError.MinBudget = &messages
			case "maxBudget":
				validationError.MaxBudget = &messages
			case "isActive":
				validationError.IsActive = &messages
			}
		}
	}

	return validationError
}
