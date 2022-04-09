package incomesource

import (
	"time"

	"github.com/emvi/null"
	"gorm.io/gorm"
)

type IncomeSource struct {
	ID        uint           `gorm:"primarykey" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      null.String    `gorm:"notnull" json:"name"`
	Balance   float64        `json:"balance"`
	User_id   null.String    `gorm:"notnull" json:"user_id"`
}
