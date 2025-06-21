package businessservices

import (
	businessapi "apps/api/business"
	models "apps/models"
	"apps/test/factories"
	"errors"
	"testing"
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestToProjectServiceSuite struct {
	WithDBSuite
}

var (
	testToProjectService ToProjectService
	company *models.Company
	project1 *models.Project
	project2 *models.Project
	project3 *models.Project
	project4 *models.Project
	project5 *models.Project
	supporter *models.Supporter
)

func (s *TestToProjectServiceSuite) SetupTest() {
	s.SetDBCon()

	testToProjectService = NewToProjectService(DBCon)

	company = factories.CompanyFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Company)
	DBCon.Create(company)

	project1 = factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 24, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 24, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project1)
	project2 = factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 25, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 25, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project2)
	project3 = factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 26, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 26, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project3)
	project4 = factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 27, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 27, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project4)
	project5 = factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 28, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 28, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project5)

	supporter = factories.SupporterFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.Supporter)
	DBCon.Create(supporter)
}

func (s *TestToProjectServiceSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestToProjectServiceSuite) TestToProjectFetchLists_EmptyArgs_NotHavingNextPage_StatusOK() {
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project1.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(0, "", "", supporter.ID)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, fetchedProducts[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[4].ProposalStatus)
	assert.Equal(s.T(), 0, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestToProjectFetchLists_EmptyArgs_HavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project6)

	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID) .Where("id != ?", project6.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project1.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(0, "", "", supporter.ID)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, fetchedProducts[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[4].ProposalStatus)
	assert.Equal(s.T(), project6.ID, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestToProjectFetchLists_WithOnlyPageToken_NotHavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ?", project2.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project4.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(project2.ID, "", "", supporter.ID)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, fetchedProducts[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[4].ProposalStatus)
	assert.Equal(s.T(), 0, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestToProjectFetchLists_WithOnlyPageToken_HavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project6)
	project7 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project7)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ?", project2.ID).Where("id < ?", project7.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project4.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(project2.ID, "", "", supporter.ID)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, fetchedProducts[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[4].ProposalStatus)
	assert.Equal(s.T(), project7.ID, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestToProjectFetchLists_WithOnlyStartDate_NotHavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ?", project2.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project4.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(0, "2025-05-25", "", supporter.ID)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, fetchedProducts[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[4].ProposalStatus)
	assert.Equal(s.T(), 0, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestToProjectFetchLists_WithOnlyStartDate_HavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project1.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(0, "2025-05-24", "", supporter.ID)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, fetchedProducts[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[4].ProposalStatus)
	assert.Equal(s.T(), project6.ID, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestToProjectFetchLists_WithOnlyEndDate_NotHavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project1.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(0, "", "2025-05-28", supporter.ID)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, fetchedProducts[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[4].ProposalStatus)
	assert.Equal(s.T(), 0, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestToProjectFetchLists_WithOnlyEndDate_HavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project1.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(0, "", "2025-05-29", supporter.ID)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, fetchedProducts[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[4].ProposalStatus)
	assert.Equal(s.T(), project6.ID, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestToProjectFetchLists_WithPageTokenAndStartDate_NotHavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project2.ID, project6.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project4.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(project2.ID, "2025-05-25", "", supporter.ID)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, fetchedProducts[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[4].ProposalStatus)
	assert.Equal(s.T(), 0, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestToProjectFetchLists_WithPageTokenAndStartDate_HavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project1.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(project1.ID, "2025-05-24", "", supporter.ID)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, fetchedProducts[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[4].ProposalStatus)
	assert.Equal(s.T(), project6.ID, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestToProjectFetchLists_WithPageTokenAndEndDate_NotHavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project2.ID, project6.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project4.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(project2.ID, "", "2025-05-29", supporter.ID)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, fetchedProducts[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[4].ProposalStatus)
	assert.Equal(s.T(), 0, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestToProjectFetchLists_WithPageTokenAndEndDate_HavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project1.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(project1.ID, "", "2025-05-29", supporter.ID)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, fetchedProducts[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[4].ProposalStatus)
	assert.Equal(s.T(), project6.ID, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestToProjectFetchLists_WithStartDateAndEndDate_NotHavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project2.ID, project6.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project4.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(0, "2025-05-25", "2025-05-29", supporter.ID)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, fetchedProducts[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[4].ProposalStatus)
	assert.Equal(s.T(), 0, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestToProjectFetchLists_WithStartDateAndEndDate_HavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project1.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(0, "2025-05-24", "2025-05-29", supporter.ID)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, fetchedProducts[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[4].ProposalStatus)
	assert.Equal(s.T(), project6.ID, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestToProjectFetchLists_WithPageTokenAndStartDateAndEndDate_NotHavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project2.ID, project6.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project4.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(project2.ID, "2025-05-25", "2025-05-29", supporter.ID)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, fetchedProducts[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[4].ProposalStatus)
	assert.Equal(s.T(), 0, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestToProjectFetchLists_WithPageTokenAndStartDateAndEndDate_HavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project1.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project2.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project3.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(project1.ID, "2025-05-24", "2025-05-29", supporter.ID)
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, fetchedProducts[0].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[1].ProposalStatus)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProducts[2].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[3].ProposalStatus)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProducts[4].ProposalStatus)
	assert.Equal(s.T(), project6.ID, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestToProjectFetch_NotProposedPlan() {
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	DBCon.Model(project).Take(project)

	fetchedProduct, err := testToProjectService.Fetch(project.ID, supporter.ID)
	
	// NOTE: レスポンスのprojectの値の確認
	assert.Equal(s.T(), project.ID, fetchedProduct.ID)
	assert.Equal(s.T(), project.Title, fetchedProduct.Title)
	assert.Equal(s.T(), businessapi.NOTPROPOSED, fetchedProduct.ProposalStatus)

	assert.Nil(s.T(), err)
}

func (s *TestToProjectServiceSuite) TestToProjectFetch_TemporaryCreatingPlan() {
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	DBCon.Model(project).Take(project)

	temporaryCreatingPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project.ID, "Status": models.PlanStatusTempraryCreating}).(*models.Plan)
	DBCon.Create(temporaryCreatingPlan)

	fetchedProduct, err := testToProjectService.Fetch(project.ID, supporter.ID)
	
	// NOTE: レスポンスのprojectの値の確認
	assert.Equal(s.T(), project.ID, fetchedProduct.ID)
	assert.Equal(s.T(), project.Title, fetchedProduct.Title)
	assert.Equal(s.T(), businessapi.TEMPORARYCREATING, fetchedProduct.ProposalStatus)

	assert.Nil(s.T(), err)
}

func (s *TestToProjectServiceSuite) TestToProjectFetch_SubmittedPlan() {
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	DBCon.Model(project).Take(project)
	
	submittedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project.ID, "Status": models.PlanStatusSubmitted}).(*models.Plan)
	DBCon.Create(submittedPlan)

	fetchedProduct, err := testToProjectService.Fetch(project.ID, supporter.ID)
	
	// NOTE: レスポンスのprojectの値の確認
	assert.Equal(s.T(), project.ID, fetchedProduct.ID)
	assert.Equal(s.T(), project.Title, fetchedProduct.Title)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProduct.ProposalStatus)

	assert.Nil(s.T(), err)
}

func (s *TestToProjectServiceSuite) TestToProjectFetch_AboveSubmittedPlan() {
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	DBCon.Model(project).Take(project)
	
	agreedPlan := factories.PlanFactory.MustCreateWithOption(map[string]interface{}{"SupporterID": supporter.ID, "ProjectID": project.ID, "Status": models.PlanStatusAgreed}).(*models.Plan)
	DBCon.Create(agreedPlan)

	fetchedProduct, err := testToProjectService.Fetch(project.ID, supporter.ID)
	
	// NOTE: レスポンスのprojectの値の確認
	assert.Equal(s.T(), project.ID, fetchedProduct.ID)
	assert.Equal(s.T(), project.Title, fetchedProduct.Title)
	assert.Equal(s.T(), businessapi.PROPOSED, fetchedProduct.ProposalStatus)

	assert.Nil(s.T(), err)
}

func (s *TestToProjectServiceSuite) TestToProjectFetch_NotFound() {
	project := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project)
	DBCon.Model(project).Take(project)

	fetchedProduct, err := testToProjectService.Fetch(project.ID+1, supporter.ID)
	
	// NOTE: レスポンスのprojectの値の確認
	assert.Equal(s.T(), 0, fetchedProduct.ID)
	assert.Equal(s.T(), "", fetchedProduct.Title)

	assert.Equal(s.T(), errors.New("not found"), err)
}

func (s *TestToProjectServiceSuite) TestCreatePlan_Success_WithoutPlanSteps() {
	startDate := time.Date(2025, 6, 1, 0, 0, 0, 0, time.Local)
	endDate := time.Date(2025, 6, 30, 0, 0, 0, 0, time.Local)

	requestParams := &businessapi.PlanStoreWithStepsInput{
		Title:       "テスト提案",
		Description: "テスト提案の概要です",
		StartDate:   &openapi_types.Date{Time: startDate},
		EndDate:     &openapi_types.Date{Time: endDate},
		UnitPrice:   5000,
		PlanSteps:   nil,
	}

	createdPlan, validationErrors, err := testToProjectService.CreatePlan(project1.ID, requestParams, supporter.ID)

	// NOTE: 成功時の確認
	assert.Nil(s.T(), validationErrors)
	assert.Nil(s.T(), err)
	assert.NotZero(s.T(), createdPlan.ID)
	assert.Equal(s.T(), supporter.ID, createdPlan.SupporterID)
	assert.Equal(s.T(), project1.ID, createdPlan.ProjectID)
	assert.Equal(s.T(), "テスト提案", createdPlan.Title)
	assert.Equal(s.T(), "テスト提案の概要です", createdPlan.Description)
	assert.Equal(s.T(), startDate, createdPlan.StartDate.Time)
	assert.Equal(s.T(), endDate, createdPlan.EndDate.Time)
	assert.Equal(s.T(), 5000, createdPlan.UnitPrice)
	assert.Equal(s.T(), models.PlanStatusTempraryCreating, createdPlan.Status)
	assert.Empty(s.T(), createdPlan.PlanSteps)

	// NOTE: データベースに保存されているかの確認
	var savedPlan models.Plan
	DBCon.First(&savedPlan, createdPlan.ID)
	assert.Equal(s.T(), createdPlan.ID, savedPlan.ID)
	assert.Equal(s.T(), supporter.ID, createdPlan.SupporterID)
	assert.Equal(s.T(), "テスト提案", savedPlan.Title)
	assert.Equal(s.T(), "テスト提案の概要です", savedPlan.Description)
	assert.Equal(s.T(), startDate, savedPlan.StartDate.Time)
	assert.Equal(s.T(), endDate, savedPlan.EndDate.Time)
	assert.Equal(s.T(), 5000, savedPlan.UnitPrice)
	assert.Equal(s.T(), models.PlanStatusTempraryCreating, savedPlan.Status)
}

func (s *TestToProjectServiceSuite) TestCreatePlan_Success_WithPlanSteps() {
	startDate := time.Date(2025, 6, 1, 0, 0, 0, 0, time.Local)
	endDate := time.Date(2025, 6, 30, 0, 0, 0, 0, time.Local)

	planSteps := []businessapi.PlanStepInput{
		{
			Title:       "ステップ1",
			Description: "ステップ1の概要",
			Duration:    10,
		},
		{
			Title:       "ステップ2",
			Description: "ステップ2の概要",
			Duration:    20,
		},
	}

	requestParams := &businessapi.PlanStoreWithStepsInput{
		Title:       "ステップ付き提案",
		Description: "ステップ付き提案の概要です",
		StartDate:   &openapi_types.Date{Time: startDate},
		EndDate:     &openapi_types.Date{Time: endDate},
		UnitPrice:   5000,
		PlanSteps:   &planSteps,
	}

	createdPlan, validationErrors, err := testToProjectService.CreatePlan(project1.ID, requestParams, supporter.ID)

	// NOTE: 成功時の確認
	assert.Nil(s.T(), validationErrors)
	assert.Nil(s.T(), err)
	assert.NotZero(s.T(), createdPlan.ID)
	assert.Equal(s.T(), supporter.ID, createdPlan.SupporterID)
	assert.Equal(s.T(), "ステップ付き提案", createdPlan.Title)
	assert.Equal(s.T(), "ステップ付き提案の概要です", createdPlan.Description)
	assert.Equal(s.T(), startDate, createdPlan.StartDate.Time)
	assert.Equal(s.T(), endDate, createdPlan.EndDate.Time)
	assert.Equal(s.T(), 5000, createdPlan.UnitPrice)
	assert.Equal(s.T(), models.PlanStatusTempraryCreating, createdPlan.Status)
	assert.Equal(s.T(), "ステップ1", createdPlan.PlanSteps[0].Title)
	assert.Equal(s.T(), "ステップ1の概要", createdPlan.PlanSteps[0].Description)
	assert.Equal(s.T(), 10, createdPlan.PlanSteps[0].Duration)
	assert.Equal(s.T(), "ステップ2", createdPlan.PlanSteps[1].Title)
	assert.Equal(s.T(), "ステップ2の概要", createdPlan.PlanSteps[1].Description)
	assert.Equal(s.T(), 20, createdPlan.PlanSteps[1].Duration)

	// NOTE: PlanStepsが保存されているかの確認
	var savedPlanSteps []models.PlanStep
	DBCon.Where("plan_id = ?", createdPlan.ID).Find(&savedPlanSteps)
	assert.Len(s.T(), savedPlanSteps, 2)
	assert.Equal(s.T(), "ステップ1", savedPlanSteps[0].Title)
	assert.Equal(s.T(), "ステップ1の概要", savedPlanSteps[0].Description)
	assert.Equal(s.T(), 10, savedPlanSteps[0].Duration)
	assert.Equal(s.T(), "ステップ2", savedPlanSteps[1].Title)
	assert.Equal(s.T(), "ステップ2の概要", savedPlanSteps[1].Description)
	assert.Equal(s.T(), 20, savedPlanSteps[1].Duration)
}

func (s *TestToProjectServiceSuite) TestCreatePlan_Success_WithoutDates() {
	requestParams := &businessapi.PlanStoreWithStepsInput{
		Title:       "日付なし提案",
		Description: "日付なし提案の概要です",
		StartDate:   nil,
		EndDate:     nil,
		UnitPrice:   3000,
		PlanSteps:   nil,
	}

	createdPlan, validationErrors, err := testToProjectService.CreatePlan(project1.ID, requestParams, supporter.ID)

	// NOTE: 成功時の確認
	assert.Nil(s.T(), validationErrors)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "日付なし提案", createdPlan.Title)
	assert.False(s.T(), createdPlan.StartDate.Valid)
	assert.False(s.T(), createdPlan.EndDate.Valid)
}

func (s *TestToProjectServiceSuite) TestCreatePlan_ValidationError_EmptyTitle() {
	requestParams := &businessapi.PlanStoreWithStepsInput{
		Title:       "",
		Description: "概要です",
		UnitPrice:   5000,
	}

	createdPlan, validationErrors, err := testToProjectService.CreatePlan(project1.ID, requestParams, supporter.ID)

	// NOTE: バリデーションエラーの確認
	assert.NotNil(s.T(), validationErrors)
	assert.Nil(s.T(), err)
	assert.Zero(s.T(), createdPlan.ID)
}

func (s *TestToProjectServiceSuite) TestCreatePlan_ValidationError_EmptyDescription() {
	requestParams := &businessapi.PlanStoreWithStepsInput{
		Title:       "テスト提案",
		Description: "",
		UnitPrice:   5000,
	}

	createdPlan, validationErrors, err := testToProjectService.CreatePlan(project1.ID, requestParams, supporter.ID)

	// NOTE: バリデーションエラーの確認
	assert.NotNil(s.T(), validationErrors)
	assert.Nil(s.T(), err)
	assert.Zero(s.T(), createdPlan.ID)
}

func (s *TestToProjectServiceSuite) TestCreatePlan_ValidationError_InvalidUnitPrice() {
	requestParams := &businessapi.PlanStoreWithStepsInput{
		Title:       "テスト提案",
		Description: "概要です",
		UnitPrice:   0,
	}

	createdPlan, validationErrors, err := testToProjectService.CreatePlan(project1.ID, requestParams, supporter.ID)

	// NOTE: バリデーションエラーの確認
	assert.NotNil(s.T(), validationErrors)
	assert.Nil(s.T(), err)
	assert.Zero(s.T(), createdPlan.ID)
}

func (s *TestToProjectServiceSuite) TestCreatePlan_ValidationError_InvalidDateRange() {
	startDate := time.Date(2025, 6, 2, 0, 0, 0, 0, time.Local)
	endDate := time.Date(2025, 6, 1, 0, 0, 0, 0, time.Local) // 開始日より前の終了日

	requestParams := &businessapi.PlanStoreWithStepsInput{
		Title:       "テスト提案",
		Description: "概要です",
		StartDate:   &openapi_types.Date{Time: startDate},
		EndDate:     &openapi_types.Date{Time: endDate},
		UnitPrice:   5000,
	}

	createdPlan, validationErrors, err := testToProjectService.CreatePlan(project1.ID, requestParams, supporter.ID)

	// NOTE: バリデーションエラーの確認
	assert.NotNil(s.T(), validationErrors)
	assert.Nil(s.T(), err)
	assert.Zero(s.T(), createdPlan.ID)
}

func (s *TestToProjectServiceSuite) TestCreatePlan_ValidationError_InvalidPlanSteps() {
	planSteps := []businessapi.PlanStepInput{
		{
			Title:       "", // 空のタイトル
			Description: "ステップ1の概要",
			Duration:    10,
		},
	}

	requestParams := &businessapi.PlanStoreWithStepsInput{
		Title:       "テスト提案",
		Description: "概要です",
		UnitPrice:   5000,
		PlanSteps:   &planSteps,
	}

	createdPlan, validationErrors, err := testToProjectService.CreatePlan(project1.ID, requestParams, supporter.ID)

	// NOTE: バリデーションエラーの確認
	assert.NotNil(s.T(), validationErrors)
	assert.Nil(s.T(), err)
	assert.Zero(s.T(), createdPlan.ID)
}

func (s *TestToProjectServiceSuite) TestCreatePlan_ProjectNotFound() {
	nonExistentProjectID := 0

	requestParams := &businessapi.PlanStoreWithStepsInput{
		Title:       "テスト提案",
		Description: "概要です",
		UnitPrice:   5000,
	}

	createdPlan, validationErrors, err := testToProjectService.CreatePlan(nonExistentProjectID, requestParams, supporter.ID)

	// NOTE: プロジェクトが見つからない場合のエラー確認
	assert.Nil(s.T(), validationErrors)
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), "project not found", err.Error())
	assert.Zero(s.T(), createdPlan.ID)
}

func (s *TestToProjectServiceSuite) TestMappingPlanWithStepsValidationErrorStruct_NoError() {
	mappedError := testToProjectService.MappingPlanWithStepsValidationErrorStruct(nil)

	// NOTE: エラーがない場合は空の構造体が返される
	assert.Nil(s.T(), mappedError.Title)
	assert.Nil(s.T(), mappedError.Description)
	assert.Nil(s.T(), mappedError.StartDate)
	assert.Nil(s.T(), mappedError.EndDate)
	assert.Nil(s.T(), mappedError.UnitPrice)
	assert.Nil(s.T(), mappedError.PlanSteps)
}

func TestToProjectService(t *testing.T) {
	// テストスイートを実行
	suite.Run(t, new(TestToProjectServiceSuite))
}
