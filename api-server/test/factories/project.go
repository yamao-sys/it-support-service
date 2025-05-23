package factories

import (
	models "apps/models"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
	"github.com/volatiletech/null/v8"
)

var ProjectFactory = factory.NewFactory(
	&models.Project{},
).Attr("CompanyID", func(args factory.Args) (interface{}, error) {
	company := CompanyFactory.MustCreate().(*models.Company)
	return company.ID, nil
}).Attr("Title", func(args factory.Args) (interface{}, error) {
	return randomdata.RandStringRunes(15), nil
}).Attr("Description", func(args factory.Args) (interface{}, error) {
	return randomdata.RandStringRunes(50), nil
}).Attr("StartDate", func(args factory.Args) (interface{}, error) {
	date := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	return date, nil
}).Attr("EndDate", func(args factory.Args) (interface{}, error) {
	date := time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)
	return date, nil
}).Attr("MinBudget", func(args factory.Args) (interface{}, error) {
	return null.Int{Int: 10000, Valid: true}, nil
}).Attr("MaxBudget", func(args factory.Args) (interface{}, error) {
	return null.Int{Int: 20000, Valid: true}, nil
}).Attr("IsActive", func(args factory.Args) (interface{}, error) {
	return true, nil
})
