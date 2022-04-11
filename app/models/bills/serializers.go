package bill

import "encoding/json"

type BillGet Bill
type BillList Bill

func (ig *BillGet) MarshalJSON() ([]byte, error) {
	type Alias BillGet
	return json.Marshal(&struct {
		ID uint `json:"id"`
		*Alias
	}{
		ig.ID,
		(*Alias)(ig),
	})
}

func (ig *BillList) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID          uint    `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Amount      float64 `json:"amount"`
	}{
		ig.ID,
		ig.Name.String,
		ig.Description.String,
		ig.Amount.Float64,
	})
}
