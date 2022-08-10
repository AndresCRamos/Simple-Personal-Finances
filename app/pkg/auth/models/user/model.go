package auth_user

import (
	bill "github.com/AndresCRamos/Simple-Personal-Finances/models/bills"
	earning "github.com/AndresCRamos/Simple-Personal-Finances/models/earning"
	incomesource "github.com/AndresCRamos/Simple-Personal-Finances/models/income_source"
	token "github.com/AndresCRamos/Simple-Personal-Finances/pkg/auth/models/token"
	"github.com/emvi/null"
)

type User struct {
	ID       uint                        `gorm:"primarykey" json:"-"`
	Name     null.String                 `gorm:"notnull" json:"name"`
	LastName null.String                 `gorm:"notnull" json:"last_name"`
	Email    null.String                 `gorm:"notnull;unique" json:"email"`
	Password []byte                      `gorm:"notnull" json:"-"`
	Token    token.Token                 `json:"-"`
	Earnings []earning.Earning           `json:"-"`
	Bills    []bill.Bill                 `json:"-"`
	Sources  []incomesource.IncomeSource `json:"-"`
}
