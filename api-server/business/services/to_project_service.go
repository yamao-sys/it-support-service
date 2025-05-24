package businessservices

import (
	models "apps/models"
	"errors"

	"gorm.io/gorm"
)

const toProjectPerPage = 5

type ToProjectService interface {
	FetchLists(pageToken int, startDate string, endDate string) (projects []models.Project, nextPageToken int)
	Fetch(ID int) (project models.Project, error error)
}

type toProjectService struct {
	db *gorm.DB
}

func NewToProjectService(db *gorm.DB) ToProjectService {
	return &toProjectService{db}
}

func (tps *toProjectService) FetchLists(pageToken int, startDate string, endDate string) (projects []models.Project, nextPageToken int) {
	query := tps.db.Model(&models.Project{})

	if pageToken > 0 {
		query = query.Where("id >= ?", pageToken)
	}
	if startDate != "" {
		query = query.Where("start_date >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("end_date <= ?", endDate)
	}
	// NOTE: nextPageTokenの検出のため、1ページの件数+1を取得
	query = query.Limit(toProjectPerPage + 1)

	query.Find(&projects)

	// NOTE: nextPageTokenのprojectをsliceから切り出し
	if len(projects) == toProjectPerPage + 1 {
		nextPageToken := projects[len(projects)-1].ID
		return projects[:len(projects)-1], nextPageToken
	}
	
	return projects, 0
}

func (tps *toProjectService) Fetch(ID int) (project models.Project, error error) {
	tps.db.First(&project, ID)
	if project.ID == 0 {
		return project, errors.New("not found")
	}
	
	return project, nil
}
