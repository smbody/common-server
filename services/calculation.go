package services

import (
	"github.com/smbody/common-server/application"
	"github.com/smbody/common-server/model"
	"github.com/smbody/common-server/scenarios/calculator"
)

func Calculate(request *application.Request) (*application.Response, error) {
	condition, err := model.NewCondition(request.Data())
	if err!=nil {return application.NewResponse(nil, err)}
	finance := calculator.NewCalculator(condition)
	calculationResult, err := finance.Calculate()
	result := calculationResult.AsSolution();
	return application.NewResponse(result, err)
}