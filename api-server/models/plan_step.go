package models

import (
	"time"
)

type PlanStep struct {
	ID          int    `gorm:"primary_key" json:"id"`
	PlanID      int    `gorm:"not null" json:"plan_id"`
	Title       string `gorm:"not null" validate:"required"`
	Description string `gorm:"not null" validate:"required"`
	Duration    int    `gorm:"not null" validate:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Plan        Plan   `gorm:"foreignKey:PlanID" validate:"omitempty"`
}
