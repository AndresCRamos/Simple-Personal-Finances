package bill

import "encoding/json"

type BillGet Bill

func (ig *Bill) MarshalJSON() ([]byte, error) {
	type Alias Bill
	return json.Marshal(&struct {
		ID uint `json:"id"`
		*Alias
	}{
		ig.ID,
		(*Alias)(ig),
	})
}
