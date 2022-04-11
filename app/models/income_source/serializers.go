package incomesource

import (
	"encoding/json"

	bill "github.com/AndresCRamos/Simple-Personal-Finances/models/bills"
	"github.com/AndresCRamos/Simple-Personal-Finances/utils"
)

type IncomeSourceGet IncomeSource
type IncomeSourceDetail IncomeSource

func getBillsById(ID uint) []bill.BillList {
	var billList []bill.Bill
	var billListDetail []bill.BillList
	utils.Instance.Find(&billList, "income_source_id = ?", ID)
	for _, currentBill := range billList {
		billListDetail = append(billListDetail, bill.BillList(currentBill))
	}
	return billListDetail
}

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
	billListDetail := getBillsById(ig.ID)
	return json.Marshal(&struct {
		*Alias
		Bills []bill.BillList `json:"bills"`
	}{
		(*Alias)(ig),
		billListDetail,
	})
}
