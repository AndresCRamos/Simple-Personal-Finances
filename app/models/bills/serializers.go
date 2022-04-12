package bill

import (
	"encoding/json"
	"fmt"
)

type BillGet Bill
type BillList Bill

func (ig *Bill) MarshalJSON() ([]byte, error) {
	year, month, day := ig.Date.Time.Date()
	return json.Marshal(&struct {
		Name             string  `json:"name"`
		Description      string  `json:"description"`
		Amount           float64 `json:"amount"`
		Date             string  `json:"date"`
		User_id          string  `json:"user_id"`
		Income_Source_id int64   `json:"source_id"`
	}{
		ig.Name.String,
		ig.Description.String,
		ig.Amount.Float64,
		fmt.Sprintf("%d-%v-%d", year, int(month), day),
		ig.User_id.String,
		ig.Income_Source_id.Int64,
	})
}

func (ig *BillGet) MarshalJSON() ([]byte, error) {
	year, month, day := ig.Date.Time.Date()
	return json.Marshal(&struct {
		ID               uint    `json:"id"`
		Name             string  `json:"name"`
		Description      string  `json:"description"`
		Amount           float64 `json:"amount"`
		Date             string  `json:"date"`
		User_id          string  `json:"user_id"`
		Income_Source_id int64   `json:"source_id"`
	}{
		ig.ID,
		ig.Name.String,
		ig.Description.String,
		ig.Amount.Float64,
		fmt.Sprintf("%d-%v-%d", year, int(month), day),
		ig.User_id.String,
		ig.Income_Source_id.Int64,
	})
}

func (ig *BillList) MarshalJSON() ([]byte, error) {
	year, month, day := ig.Date.Time.Date()
	return json.Marshal(&struct {
		ID          uint    `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Amount      float64 `json:"amount"`
		Date        string  `json:"date"`
	}{
		ig.ID,
		ig.Name.String,
		ig.Description.String,
		ig.Amount.Float64,
		fmt.Sprintf("%d-%v-%d", year, int(month), day),
	})
}
