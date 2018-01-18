package calculator

import (
	"github.com/shopspring/decimal"
	"github.com/smbody/common-server/scenarios/tools"
	"github.com/smbody/common-server/model"
)

type Overall struct {
	iterators []*toolIterator
	count int
	value decimal.Decimal
	values []decimal.Decimal

}

func NewOverall(distribution []tools.Distribution) *Overall{
	o := &Overall{}
	for _, d := range distribution {
		o.iterators = append(o.iterators, NewToolIterator(&d))
	}
	return o
}


func (o *Overall) Iteration() {
	var amount decimal.Decimal
	for _, i := range o.iterators {
		iteration := i.AddIteration();
		amount = amount.Add(iteration)
	}
	o.count++
	o.value = amount
	o.values = append(o.values, o.value);
}

func (o *Overall) Count() int {
	return o.count
}

func (o *Overall) Value() decimal.Decimal {
	return o.value
}

func (o *Overall) Values() []decimal.Decimal {
	return o.values
}

func (o *Overall) AsSolution() *model.Solution {
	result := model.NewSolution()
	result.Count = o.Count()
	result.Value = o.Value()
	result.Values = o.Values()
	for _,tool := range o.iterators {
		result.Tools = append(result.Tools, createDetail(tool))
	}

	return result

}

func createDetail(iterator *toolIterator) *model.Detail{
	result := model.NewDetail()
	result.Name = iterator.tool.Id()
	result.Percent = iterator.percent
	result.Investment = iterator.invest
	result.Value = iterator.value
	result.Values = iterator.values
	return result
}