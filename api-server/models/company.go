package models

import "time"

type Company struct {
	ID             int       `gorm:"primary_key" json:"id"`
	Name           string    `gorm:"not null" validate:"required"`
	Email          string    `gorm:"not null;unique" validate:"required,email"`
	Password       string    `gorm:"not null" validate:"required"`
	FinalTaxReturn string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
