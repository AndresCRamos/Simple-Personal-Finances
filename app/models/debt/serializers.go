package debt

import (
	"time"

	"github.com/emvi/null"
)

type DebtCreate struct {
	ID          uint         `json:"-"`
	Name        null.String  `json:"name"`
	Description null.String  `json:"description"`
	Amount      null.Float64 `json:"amount"`
	Date        null.String  `json:"date"`
	Debtor      null.Int64   `json:"debtor_id"`
}

func (bc *DebtCreate) Parse() *Debt {
	date, parsed := time.Parse("2006-01-02", bc.Date.String)
	validDate := bc.Date.Valid && bc.Date.String != "" && parsed == nil
	return &Debt{
		ID:             bc.ID,
		Name:           bc.Name,
		Description:    bc.Description,
		Amount:         bc.Amount,
		Date:           null.NewTime(date, validDate),
		Debtor_user_id: bc.Debtor,
	}
}
