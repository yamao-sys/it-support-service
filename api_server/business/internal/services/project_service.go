package services

import (
	"business/api/generated/projects"
	validator "business/internal/validators"
	models "business/models/generated"
	"context"
	"database/sql"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ProjectService interface {
	Create(ctx context.Context, requestParams *projects.PostProjectsJSONRequestBody, companyID int) (project models.Project, validatorErrors error, error error)
	MappingValidationErrorStruct(err error) projects.ProjectValidationError
}

type projectService struct {
	db *sql.DB
}

func NewProjectService(db *sql.DB) ProjectService {
	return &projectService{db}
}

func (ps *projectService) Create(ctx context.Context, requestParams *projects.PostProjectsJSONRequestBody, companyID int) (project models.Project, validatorErrors error, error error) {
	// NOTE: バリデーションチェック
	validatorErrors = validator.ValidateProject(requestParams)
	if validatorErrors != nil {
		return models.Project{}, validatorErrors, nil
	}

	project = models.Project{}
	project.CompanyID = companyID
	project.Title = *requestParams.Title
	project.Description = *requestParams.Description
	project.StartDate = requestParams.StartDate.Time
	project.EndDate = requestParams.EndDate.Time
	project.MinBudget = null.Int{Int: *requestParams.MinBudget, Valid: true}
	project.MaxBudget = null.Int{Int: *requestParams.MaxBudget, Valid: true}
	project.IsActive = *requestParams.IsActive

	createErr := project.Insert(ctx, ps.db, boil.Infer())
	if createErr != nil {
		return models.Project{}, nil, createErr
	}

	return project, nil, nil
}

func (ps *projectService) MappingValidationErrorStruct(err error) projects.ProjectValidationError {
	var validationError projects.ProjectValidationError
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
