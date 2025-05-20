package businessservices

import (
	businessapi "apps/api/business"
	businessvalidators "apps/business/validators"
	models "apps/models"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/volatiletech/null/v8"
	"gorm.io/gorm"
)

const perPage = 5

type ProjectService interface {
	FetchLists(companyID int, pageToken int) (projects []models.Project, nextPageToken int, error error)
	Create(requestParams *businessapi.ProjectStoreInput, companyID int) (project models.Project, validatorErrors error, error error)
	Fetch(ID int) (project models.Project, error error)
	Update(requestParams *businessapi.ProjectStoreInput, ID int) (project models.Project, validatorErrors error, error error)
	MappingValidationErrorStruct(err error) businessapi.ProjectValidationError
}

type projectService struct {
	db *gorm.DB
}

func NewProjectService(db *gorm.DB) ProjectService {
	return &projectService{db}
}

func (ps *projectService) FetchLists(companyID int, pageToken int) (projects []models.Project, nextPageToken int, error error) {
	query := ps.db.Model(&models.Project{})

	query = query.Where("company_id = ?", companyID)
	if pageToken > 0 {
		query = query.Where("id >= ?", pageToken)
	}
	// NOTE: nextPageTokenの検出のため、1ページの件数+1を取得
	query = query.Limit(perPage + 1)

	query.Find(&projects)

	// NOTE: nextPageTokenのprojectをsliceから切り出し
	if len(projects) == perPage + 1 {
		nextPageToken := projects[len(projects)-1].ID
		return projects[:len(projects)-1], nextPageToken, nil
	}
	
	return projects, 0, nil
}

func (ps *projectService) Create(requestParams *businessapi.ProjectStoreInput, companyID int) (project models.Project, validatorErrors error, error error) {
	// NOTE: バリデーションチェック
	validatorErrors = businessvalidators.ValidateProject(requestParams)
	if validatorErrors != nil {
		return models.Project{}, validatorErrors, nil
	}

	project = models.Project{}
	project.CompanyID = companyID
	project.Title = requestParams.Title
	project.Description = requestParams.Description
	project.StartDate = requestParams.StartDate.Time
	project.EndDate = requestParams.EndDate.Time
	if requestParams.MinBudget != nil {
		project.MinBudget = null.Int{Int: *requestParams.MinBudget, Valid: true}
	}
	if requestParams.MaxBudget != nil {
		project.MaxBudget = null.Int{Int: *requestParams.MaxBudget, Valid: true}
	}
	project.IsActive = requestParams.IsActive

	ps.db.Create(&project)

	return project, nil, nil
}

func (ps *projectService) Fetch(ID int) (project models.Project, error error) {
	ps.db.Where("id = ?", ID).Take(&project)
	if project.ID == 0 {
		return project, errors.New("not found")
	}
	
	return project, nil
}

func (ps *projectService) Update(requestParams *businessapi.ProjectStoreInput, ID int) (project models.Project, validatorErrors error, error error) {
	// NOTE: バリデーションチェック
	validatorErrors = businessvalidators.ValidateProject(requestParams)
	if validatorErrors != nil {
		return models.Project{}, validatorErrors, nil
	}

	ps.db.Where("id = ?", ID).Take(&project)
	if project.ID == 0 {
		return models.Project{}, nil, errors.New("not found")
	}
	project.Title = requestParams.Title
	project.Description = requestParams.Description
	project.StartDate = requestParams.StartDate.Time
	project.EndDate = requestParams.EndDate.Time
	if requestParams.MinBudget != nil {
		project.MinBudget = null.Int{Int: *requestParams.MinBudget, Valid: true}
	}
	if requestParams.MaxBudget != nil {
		project.MaxBudget = null.Int{Int: *requestParams.MaxBudget, Valid: true}
	}
	project.IsActive = requestParams.IsActive

	ps.db.Save(&project)

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
