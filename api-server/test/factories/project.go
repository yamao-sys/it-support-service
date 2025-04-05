package factories

import (
	models "apps/models/generated"
	"database/sql"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var ProjectFactory = factory.NewFactory(
	&models.Project{
		Title: randomdata.RandStringRunes(15),
		Description: randomdata.RandStringRunes(50),
		MinBudget: null.Int{Int: 10000, Valid: true},
		MaxBudget: null.Int{Int: 20000, Valid: true},
		IsActive: true,
	},
).Attr("CompanyID", func(args factory.Args) (interface{}, error) {
	company := CompanyFactory.MustCreate().(*models.Company)
	company.Insert(args.Context(), args.Instance().(*sql.DB), boil.Infer())
	return company.ID, nil
}).Attr("StartDate", func(args factory.Args) (interface{}, error) {
	date := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	return date, nil
}).Attr("EndDate", func(args factory.Args) (interface{}, error) {
	date := time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC)
	return date, nil
})
