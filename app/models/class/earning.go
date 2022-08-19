package class

import (
	"encoding/json"
	"fmt"
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
	Date             null.Time      `gorm:"type:date;notnull" json:"date"`
	User_id          null.Int64     `gorm:"notnull" json:"user_id"`
	Income_Source_id null.Int64     `gorm:"notnull" json:"source_id"`
}

func (ig *Earning) MarshalJSON() ([]byte, error) {
	year, month, day := ig.Date.Time.Date()
	return json.Marshal(&struct {
		Name             string  `json:"name"`
		Description      string  `json:"description"`
		Amount           float64 `json:"amount"`
		Date             string  `json:"date"`
		User_id          int64   `json:"user_id"`
		Income_Source_id int64   `json:"source_id"`
	}{
		ig.Name.String,
		ig.Description.String,
		ig.Amount.Float64,
		fmt.Sprintf("%d-%v-%d", year, int(month), day),
		ig.User_id.Int64,
		ig.Income_Source_id.Int64,
	})
}
