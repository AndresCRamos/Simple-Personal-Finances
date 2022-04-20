package earning

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/emvi/null"
)

type EarningGet Earning
type EarningList Earning

func (ig *Earning) MarshalJSON() ([]byte, error) {
	year, month, day := ig.Date.Time.Date()
	return json.Marshal(&struct {
		Name             string  `json:"name"`
		Description      string  `json:"description"`
		Amount           float64 `json:"amount"`
		Date             string  `json:"date"`
		User_id          int64   `json:"user_id"`
		Income_Source_id int64   `json:"source_id"`
	}{
		ig.Name.String,
		ig.Description.String,
		ig.Amount.Float64,
		fmt.Sprintf("%d-%v-%d", year, int(month), day),
		ig.User_id.Int64,
		ig.Income_Source_id.Int64,
	})
}

func (ig *EarningGet) MarshalJSON() ([]byte, error) {
	year, month, day := ig.Date.Time.Date()
	return json.Marshal(&struct {
		ID               uint    `json:"id"`
		Name             string  `json:"name"`
		Description      string  `json:"description"`
		Amount           float64 `json:"amount"`
		Date             string  `json:"date"`
		User_id          int64   `json:"user_id"`
		Income_Source_id int64   `json:"source_id"`
	}{
		ig.ID,
		ig.Name.String,
		ig.Description.String,
		ig.Amount.Float64,
		fmt.Sprintf("%d-%v-%d", year, int(month), day),
		ig.User_id.Int64,
		ig.Income_Source_id.Int64,
	})
}

func (ig *EarningList) MarshalJSON() ([]byte, error) {
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

type EarningCreate struct {
	ID               uint         `json:"-"`
	Name             null.String  `json:"name"`
	Description      null.String  `json:"description"`
	Amount           null.Float64 `json:"amount"`
	Date             null.String  `json:"date"`
	User_id          null.Int64   `json:"user_id"`
	Income_Source_id null.Int64   `json:"source_id"`
}

func (bc *EarningCreate) Parse() *Earning {
	date, _ := time.Parse("2006-01-02", bc.Date.String)
	validDate := bc.Date.Valid && bc.Date.String != ""
	return &Earning{
		ID:               bc.ID,
		Name:             bc.Name,
		Description:      bc.Description,
		Amount:           bc.Amount,
		User_id:          bc.User_id,
		Income_Source_id: bc.Income_Source_id,
		Date:             null.NewTime(date, validDate),
	}
}
