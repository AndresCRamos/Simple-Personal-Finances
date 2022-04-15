package earning

import "encoding/json"

type EarningGet Earning
type EarningList Earning

func (ig *EarningGet) MarshalJSON() ([]byte, error) {
	type Alias EarningGet
	return json.Marshal(&struct {
		ID uint `json:"id"`
		*Alias
	}{
		ig.ID,
		(*Alias)(ig),
	})
}

func (ig *EarningList) MarshalJSON() ([]byte, error) {
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
