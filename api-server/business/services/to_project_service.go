package businessservices

import (
	businessapi "apps/api/business"
	models "apps/models"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

const toProjectPerPage = 5

type ToProjectService interface {
	FetchLists(pageToken int, startDate string, endDate string, supporterID int) (toProjects []ToProjectFields, nextPageToken int)
	Fetch(ID int, supporterID int) (toProject ToProjectFields, error error)
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
