package models

import (
	"time"

	"github.com/volatiletech/null/v8"
)

type Plan struct {
	ID                 int       `gorm:"primary_key" json:"id"`
	SupporterID        int    `gorm:"not null" json:"supporter_id"`
	ProjectID          int    `gorm:"not null" json:"project_id"`
	Title              string    `gorm:"not null" validate:"required"`
	Description        string    `gorm:"not null" validate:"required"`
	StartDate          null.Time `gorm:"type:date;column:start_date" validate:"omitempty"`
	EndDate            null.Time `gorm:"type:date;column:end_date" validate:"omitempty"`
	UnitPrice		   int  `gorm:"column:unit_price" validate:"required"`
	Status			   int       `gorm:"not null" validate:"required"`
	AgreedAt           null.Time `gorm:"type:date;column:agreed_at" validate:"omitempty"`
	CreatedAt          time.Time
	UpdatedAt          time.Time

	Supporter          Supporter   `gorm:"foreignKey:SupporterID" validate:"omitempty"`
	Project            Project   `gorm:"foreignKey:ProjectID" validate:"omitempty"`
}
