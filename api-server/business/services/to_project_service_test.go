package businessservices

import (
	models "apps/models"
	"apps/test/factories"
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
}

func (s *TestToProjectServiceSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestToProjectServiceSuite) TestProjectFetchLists_EmptyArgs_NotHavingNextPage_StatusOK() {
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Pluck("id", &companyProductIDs)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(0, "", "")
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), 0, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestProjectFetchLists_EmptyArgs_HavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project6)

	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID) .Where("id != ?", project6.ID).Pluck("id", &companyProductIDs)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(0, "", "")
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), project6.ID, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestProjectFetchLists_WithOnlyPageToken_NotHavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ?", project2.ID).Pluck("id", &companyProductIDs)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(project2.ID, "", "")
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), 0, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestProjectFetchLists_WithOnlyPageToken_HavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project6)
	project7 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID}).(*models.Project)
	DBCon.Create(project7)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ?", project2.ID).Where("id < ?", project7.ID).Pluck("id", &companyProductIDs)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(project2.ID, "", "")
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), project7.ID, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestProjectFetchLists_WithOnlyStartDate_NotHavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ?", project2.ID).Pluck("id", &companyProductIDs)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(0, "2025-05-25", "")
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), 0, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestProjectFetchLists_WithOnlyStartDate_HavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(0, "2025-05-24", "")
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), project6.ID, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestProjectFetchLists_WithOnlyEndDate_NotHavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(0, "", "2025-05-28")
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), 0, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestProjectFetchLists_WithOnlyEndDate_HavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(0, "", "2025-05-29")
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), project6.ID, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestProjectFetchLists_WithPageTokenAndStartDate_NotHavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project2.ID, project6.ID).Pluck("id", &companyProductIDs)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(project2.ID, "2025-05-25", "")
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), 0, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestProjectFetchLists_WithPageTokenAndStartDate_HavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(project1.ID, "2025-05-24", "")
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), project6.ID, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestProjectFetchLists_WithPageTokenAndEndDate_NotHavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project2.ID, project6.ID).Pluck("id", &companyProductIDs)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(project2.ID, "", "2025-05-29")
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), 0, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestProjectFetchLists_WithPageTokenAndEndDate_HavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(project1.ID, "", "2025-05-29")
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), project6.ID, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestProjectFetchLists_WithStartDateAndEndDate_NotHavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project2.ID, project6.ID).Pluck("id", &companyProductIDs)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(0, "2025-05-25", "2025-05-29")
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), 0, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestProjectFetchLists_WithStartDateAndEndDate_HavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(0, "2025-05-24", "2025-05-29")
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), project6.ID, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestProjectFetchLists_WithPageTokenAndStartDateAndEndDate_NotHavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project2.ID, project6.ID).Pluck("id", &companyProductIDs)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(project2.ID, "2025-05-25", "2025-05-29")
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), 0, nextPageToken)
}

func (s *TestToProjectServiceSuite) TestProjectFetchLists_WithPageTokenAndStartDateAndEndDate_HavingNextPage_StatusOK() {
	project6 := factories.ProjectFactory.MustCreateWithOption(map[string]interface{}{"CompanyID": company.ID, "StartDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local), "EndDate": time.Date(2025, 5, 29, 0, 0, 0, 0, time.Local)}).(*models.Project)
	DBCon.Create(project6)
	var companyProductIDs []int
	DBCon.Model(&models.Project{}).Where("company_id = ?", company.ID).Where("id >= ? AND id <= ?", project1.ID, project5.ID).Pluck("id", &companyProductIDs)

	fetchedProducts, nextPageToken := testToProjectService.FetchLists(project1.ID, "2025-05-24", "2025-05-29")
	var fetchedProductIDs []int
	for _, fetchedProduct := range fetchedProducts {
		fetchedProductIDs = append(fetchedProductIDs, fetchedProduct.ID)
	}
	assert.ElementsMatch(s.T(), companyProductIDs, fetchedProductIDs)
	assert.Equal(s.T(), project6.ID, nextPageToken)
}

func TestToProjectService(t *testing.T) {
	// テストスイートを実行
	suite.Run(t, new(TestToProjectServiceSuite))
}
