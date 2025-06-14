package businessservices

import (
	businessapi "apps/api/business"
	models "apps/models"
	"apps/test/factories"
	"errors"
	"testing"
	"time"

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

func TestToProjectService(t *testing.T) {
	// テストスイートを実行
	suite.Run(t, new(TestToProjectServiceSuite))
}
