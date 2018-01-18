package tools

import (
	"github.com/shopspring/decimal"
	"github.com/smbody/common-server/model"
	"github.com/smbody/common-server/application"
)

type Distribution struct {
	Tool    model.Tool
	Percent decimal.Decimal
	Invest  decimal.Decimal
}

func NewDistribution(risk int, profit decimal.Decimal) []Distribution {
	//todo: impl
	return []Distribution{
		//Distribution{NewPiggyBank(), decimal.NewFromFloat(100), profit}}
	Distribution{NewPiggyBank(), application.IntToDecimal(20), profit.Mul(application.StrToDecimal("0.2"))},
	Distribution{NewDeposit(application.StrToDecimal("6.86")), application.IntToDecimal(80), profit.Mul(application.StrToDecimal("0.8"))}}
}
