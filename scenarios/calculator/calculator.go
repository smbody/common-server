package calculator

import (
	"fmt"

	"github.com/smbody/common-server/model"
	"github.com/smbody/common-server/scenarios/tools"
)

type Calculator struct {
	condition *model.Condition
	overall   *Overall
}

func NewCalculator(condition *model.Condition) *Calculator {
	c := &Calculator{}
	c.condition = condition
	c.overall = NewOverall(tools.NewDistribution(condition.Risk, condition.Incomes.Sub(condition.Outcomes)))
	return c
}

func (c *Calculator) Calculate() (*Overall, error) {
	for i := 0; i < 6000; i++ {
		if c.overall.Iteration(); c.overall.Value().Cmp(c.condition.Target) >=0  {
			return c.overall, nil
		}
	}
	return c.overall, fmt.Errorf("Can't get target after %v iterations.", c.overall.Count)
}
