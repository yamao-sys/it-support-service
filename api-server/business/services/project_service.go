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

type ProjectService interface {
	FetchLists(ctx context.Context, companyID int) (projects models.ProjectSlice, error error)
	Create(ctx context.Context, requestParams *businessapi.ProjectStoreInput, companyID int) (project models.Project, validatorErrors error, error error)
	Update(ctx context.Context, requestParams *businessapi.ProjectStoreInput, ID int) (project models.Project, validatorErrors error, error error)
	MappingValidationErrorStruct(err error) businessapi.ProjectValidationError
}

type projectService struct {
	db *sql.DB
}

func NewProjectService(db *sql.DB) ProjectService {
	return &projectService{db}
}

func (ps *projectService) FetchLists(ctx context.Context, companyID int) (projects models.ProjectSlice, error error) {
	return models.Projects(qm.Where("company_id = ?", companyID)).All(ctx, ps.db)
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
	project.MinBudget = null.Int{Int: *requestParams.MinBudget, Valid: true}
	project.MaxBudget = null.Int{Int: *requestParams.MaxBudget, Valid: true}
	project.IsActive = *requestParams.IsActive

	createErr := project.Insert(ctx, ps.db, boil.Infer())
	if createErr != nil {
		return models.Project{}, nil, createErr
	}

	return project, nil, nil
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
	doUpdateProject.MinBudget = null.Int{Int: *requestParams.MinBudget, Valid: true}
	doUpdateProject.MaxBudget = null.Int{Int: *requestParams.MaxBudget, Valid: true}
	doUpdateProject.IsActive = *requestParams.IsActive

	_, updateErr := doUpdateProject.Update(ctx, ps.db, boil.Infer())
	if updateErr != nil {
		return models.Project{}, nil, updateErr
	}

	return *doUpdateProject, nil, nil
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
