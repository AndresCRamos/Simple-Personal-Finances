package incomesource

import (
	"encoding/json"

	bill "github.com/AndresCRamos/Simple-Personal-Finances/models/bills"
)

type IncomeSourceGet IncomeSource
type IncomeSourceDetail IncomeSource

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
	billListDetail := bill.GetBillsBySourceId(ig.ID)
	return json.Marshal(&struct {
		*Alias
		Bills []bill.BillList `json:"bills"`
	}{
		(*Alias)(ig),
		billListDetail,
	})
}
