package application

import "github.com/shopspring/decimal"

func init() {
	decimal.DivisionPrecision = 2
}

func StrToDecimal(source string) decimal.Decimal {
	result, err := decimal.NewFromString(source)
	if err != nil {result = decimal.Zero}
	return result
}

func IntToDecimal(source int) decimal.Decimal {
	return decimal.New(int64(source), 0)
}
