package model

import "github.com/shopspring/decimal"

type Solution struct {
	Count int
	Value decimal.Decimal
	Values []decimal.Decimal
	Tools []*Detail
}

type Detail struct {
	Name string
	Percent decimal.Decimal
	Investment decimal.Decimal
	Value decimal.Decimal
	Values []decimal.Decimal
}

func NewSolution() *Solution {
	return &Solution{}
}

func NewDetail() *Detail {
	return &Detail{}
}