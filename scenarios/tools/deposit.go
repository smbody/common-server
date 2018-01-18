package tools

import "github.com/shopspring/decimal"

type deposit struct {
	percent         decimal.Decimal
	daysInIteration decimal.Decimal
	daysInYear      decimal.Decimal
}

func NewDeposit(percent decimal.Decimal) *deposit {
	result := &deposit{}
	result.percent = percent
	result.daysInYear, _ = decimal.NewFromString("365")
	result.daysInIteration, _ = decimal.NewFromString("30")
	return result
}

func (p *deposit) Id() string {
	return "deposit"
}

func (p *deposit) Calculate(startFrom decimal.Decimal, invest decimal.Decimal) decimal.Decimal {
	amount := startFrom.Add(invest)
	percentAmount := amount.Mul(p.percent.Mul(p.daysInIteration)).Div(p.daysInYear).Div(decimal.NewFromFloat(100))
	return amount.Add(percentAmount)
}
