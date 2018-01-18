package model

import "github.com/shopspring/decimal"

type Tool interface {
	Id() string
	Calculate(startFrom decimal.Decimal, invest decimal.Decimal) decimal.Decimal
}


