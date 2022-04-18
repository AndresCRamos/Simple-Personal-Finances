package bill

import (
	"time"

	"github.com/emvi/null"
	"gorm.io/gorm"
)

type Bill struct {
	ID               uint           `gorm:"primarykey" json:"-"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	Name             null.String    `gorm:"notnull" json:"name"`
	Description      null.String    `json:"description"`
	Amount           null.Float64   `json:"amount"`
	Date             null.Time      `gorm:"type:date;notnull" json:"date"`
	User_id          null.Int64     `gorm:"notnull" json:"user_id"`
	Income_Source_id null.Int64     `gorm:"notnull" json:"source_id"`
}
