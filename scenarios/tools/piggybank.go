package tools

import "github.com/shopspring/decimal"

type piggyBank struct {
}

func NewPiggyBank() *piggyBank {
	return &piggyBank{}
}

func (p *piggyBank) Id() string {
	return "piggybank"
}

func (p *piggyBank) Calculate(startFrom decimal.Decimal, invest decimal.Decimal) decimal.Decimal {
	return startFrom.Add(invest)
}
