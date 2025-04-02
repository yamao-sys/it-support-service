package factories

import (
	models "apps/models/generated"
	"log"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
	"golang.org/x/crypto/bcrypt"
)

var SupporterFactory = factory.NewFactory(
	&models.Supporter{
		FirstName:  randomdata.FirstName(randomdata.RandomGender),
		LastName: randomdata.LastName(),
		Email: randomdata.Email(),
		FrontIdentification: randomdata.StringSample(),
		BackIdentification: randomdata.StringSample(),
	},
).Attr("Password", func(args factory.Args) (interface{}, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to generate hash %v", err)
	}
	return string(hash), nil
})
