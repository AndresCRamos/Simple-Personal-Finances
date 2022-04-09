package incomesource

import "encoding/json"

type IncomeSourceGet IncomeSource

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
