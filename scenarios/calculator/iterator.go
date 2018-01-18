package calculator

import (
	"github.com/shopspring/decimal"
	"github.com/smbody/common-server/model"
	"github.com/smbody/common-server/scenarios/tools"
)

type toolIterator struct {
	tool    model.Tool
	percent decimal.Decimal
	invest  decimal.Decimal
	value   decimal.Decimal
	values []decimal.Decimal
}

func NewToolIterator(d *tools.Distribution) *toolIterator {
	return &toolIterator{tool: d.Tool, percent: d.Percent, invest: d.Invest, value:decimal.Zero}
}

func (i *toolIterator) AddIteration() decimal.Decimal {
	i.value = i.tool.Calculate(i.value, i.invest)
	i.values = append(i.values, i.value)
	return i.value
}

