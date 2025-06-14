package factories

import (
	models "apps/models"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
	"github.com/volatiletech/null/v8"
)

var PlanFactory = factory.NewFactory(
	&models.Plan{},
).Attr("SupporterID", func(args factory.Args) (interface{}, error) {
	supporter := SupporterFactory.MustCreate().(*models.Supporter)
	return supporter.ID, nil
}).Attr("ProjectID", func(args factory.Args) (interface{}, error) {
	project := ProjectFactory.MustCreate().(*models.Project)
	return project.ID, nil
}).Attr("Title", func(args factory.Args) (interface{}, error) {
	return randomdata.RandStringRunes(15), nil
}).Attr("Description", func(args factory.Args) (interface{}, error) {
	return randomdata.RandStringRunes(50), nil
}).Attr("StartDate", func(args factory.Args) (interface{}, error) {
	date := null.Time{Time: time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC), Valid: true}
	return date, nil
}).Attr("EndDate", func(args factory.Args) (interface{}, error) {
	date := null.Time{Time: time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC), Valid: true}
	return date, nil
}).Attr("UnitPrice", func(args factory.Args) (interface{}, error) {
	return 10000, nil
}).Attr("Status", func(args factory.Args) (interface{}, error) {
	return models.PlanStatusNum(models.PlanStatusTempraryCreating), nil
}).Attr("AgreedAt", func(args factory.Args) (interface{}, error) {
	date := null.Time{Time: time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC), Valid: true}
	return date, nil
})
