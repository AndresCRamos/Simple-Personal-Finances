package incomesource

import (
	"encoding/json"

	bill "github.com/AndresCRamos/Simple-Personal-Finances/models/bills"
	"github.com/AndresCRamos/Simple-Personal-Finances/models/class"
	earning "github.com/AndresCRamos/Simple-Personal-Finances/models/earning"
)

type IncomeSourceGet class.IncomeSource
type IncomeSourceDetail class.IncomeSource

func (ig *IncomeSourceGet) MarshalJSON() ([]byte, error) {
	type Alias IncomeSourceGet
	return json.Marshal(&struct {
		ID uint `json:"id"`
		*Alias
	}{
		ig.ID,
		(*Alias)(ig),
	})
}

func (ig *IncomeSourceDetail) MarshalJSON() ([]byte, error) {
	type Alias IncomeSourceDetail
	billListDetail := bill.GetBillsBySourceId(ig.ID, uint(ig.User_id.Int64))
	earningListDetail := earning.GetEarningsBySourceId(ig.ID, uint(ig.User_id.Int64))
	return json.Marshal(&struct {
		*Alias
		Bills    []bill.BillList       `json:"bills"`
		Earnings []earning.EarningList `json:"earnings"`
	}{
		(*Alias)(ig),
		billListDetail,
		earningListDetail,
	})
}
