package earning

import (
	"time"

	"github.com/emvi/null"
	"gorm.io/gorm"
)

type Earning struct {
	ID               uint           `gorm:"primarykey" json:"-"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	Name             null.String    `gorm:"notnull" json:"name"`
	Description      null.String    `json:"description"`
	Amount           null.Float64   `json:"amount"`
	User_id          null.String    `gorm:"notnull" json:"user_id"`
	Income_Source_id null.Int64     `gorm:"notnull" json:"source_id"`
}
