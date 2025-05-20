package models

import (
	"time"

	"github.com/volatiletech/null/v8"
)

type Supporter struct {
	ID                  int       `gorm:"primary_key" json:"id"`
	FirstName           string    `gorm:"not null;column:first_name" validate:"required"`
	LastName            string    `gorm:"not null;column:last_name" validate:"required"`
	Email               string    `gorm:"not null;unique" validate:"required,email"`
	Password            string    `gorm:"not null" validate:"required"`
	Birthday            null.Time `gorm:"type:date" validate:"omitempty"`
	FrontIdentification string    `gorm:"column:front_identification" validate:"required"`
	BackIdentification  string    `gorm:"column:back_identification" validate:"required"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
