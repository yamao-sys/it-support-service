package models

import (
	"time"

	"github.com/volatiletech/null/v8"
)

type Project struct {
	ID                 int       `gorm:"primary_key" json:"id"`
	CompanyID          int    `gorm:"not null" json:"company_id"`
	Title              string    `gorm:"not null" validate:"required"`
	Description        string    `gorm:"not null" validate:"required"`
	StartDate          time.Time `gorm:"type:date;column:start_date" validate:"omitempty"`
	EndDate            time.Time `gorm:"type:date;column:end_date" validate:"omitempty"`
	MinBudget		   null.Int  `gorm:"column:min_budget" validate:"omitempty"`
	MaxBudget		   null.Int  `gorm:"column:max_budget" validate:"omitempty"`
	IsActive		   bool      `gorm:"not null" validate:"required"`
	CreatedAt          time.Time
	UpdatedAt          time.Time

	Company            Company   `gorm:"foreignKey:CompanyID" validate:"omitempty"`
	Plans              []Plan    `gorm:"foreignKey:ProjectID" validate:"omitempty"`
}
