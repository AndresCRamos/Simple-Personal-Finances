package incomesource

import "gorm.io/gorm"

type IncomeSource struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Balance float64
	User_id string `gorm:"not null"`
}
