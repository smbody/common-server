package model

import (
	"io"
	"encoding/json"
	"github.com/shopspring/decimal"
)

type Condition struct {
	Incomes  decimal.Decimal
	Outcomes decimal.Decimal
	Target   decimal.Decimal
	Risk     int
}

func NewCondition(reader io.Reader) (result *Condition, err error) {
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&result)
	return
}
