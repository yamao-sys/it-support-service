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

var PlanFactory = factory.NewFactory(
	&models.Plan{},
).Attr("SupporterID", func(args factory.Args) (interface{}, error) {
	supporter := SupporterFactory.MustCreate().(*models.Supporter)
	supporter.Insert(args.Context(), args.Instance().(*sql.DB), boil.Infer())
	return supporter.ID, nil
}).Attr("ProjectID", func(args factory.Args) (interface{}, error) {
	project := ProjectFactory.MustCreate().(*models.Project)
	project.Insert(args.Context(), args.Instance().(*sql.DB), boil.Infer())
	return project.ID, nil
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
}).Attr("UnitPrice", func(args factory.Args) (interface{}, error) {
	return null.Int{Int: 10000, Valid: true}, nil
}).Attr("Status", func(args factory.Args) (interface{}, error) {
	return 0, nil
}).Attr("AgreedAt", func(args factory.Args) (interface{}, error) {
	date := time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)
	return date, nil
})
