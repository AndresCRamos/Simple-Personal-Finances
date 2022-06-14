package debt

import (
	"time"

	"github.com/emvi/null"
	"gorm.io/gorm"
)

type Debt struct {
	ID             uint           `gorm:"primarykey" json:"-"`
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	Name           null.String    `gorm:"notnull" json:"name"`
	Description    null.String    `json:"description"`
	Amount         null.Float64   `json:"amount"`
	Date           null.Time      `gorm:"type:date;notnull" json:"date"`
	Lender_user_id null.Int64     `gorm:"notnull" json:"-"`
	Debtor_user_id null.Int64     `gorm:"notnull" json:"-"`
}
