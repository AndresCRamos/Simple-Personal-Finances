package auth_user

import (
	"github.com/AndresCRamos/Simple-Personal-Finances/models/class"
	token "github.com/AndresCRamos/Simple-Personal-Finances/pkg/auth/models/token"
	"github.com/emvi/null"
)

type User struct {
	ID       uint                 `gorm:"primarykey" json:"-"`
	Name     null.String          `gorm:"notnull" json:"name"`
	LastName null.String          `gorm:"notnull" json:"last_name"`
	Email    null.String          `gorm:"notnull;unique" json:"email"`
	Password []byte               `gorm:"notnull" json:"-"`
	Token    token.Token          `json:"-"`
	Earnings []class.Earning      `json:"-"`
	Bills    []class.Bill         `json:"-"`
	Sources  []class.IncomeSource `json:"-"`
}
