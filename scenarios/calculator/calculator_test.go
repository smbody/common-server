package calculator

import (
	"testing"

	"github.com/smbody/common-server/model"
	"github.com/smbody/common-server/application"
)

func TestCalculate(t *testing.T) {
	t.Log("Passed conditions affect the calculation.")
	condition := &model.Condition{Incomes: application.StrToDecimal("1000"), Outcomes: application.StrToDecimal("890"), Target: application.StrToDecimal("3000")}
	c := NewCalculator(condition)
	if result, err := c.Calculate(); result.Value().Cmp(condition.Target) < 0 {
		t.Errorf("Expected result: '%v' but was '%v'. With error -> %s", condition.Target, result.Value, err)
		t.Logf("Result '%v' after '%v' iterations.", result.Value, result.Count)
		t.Logf("Overall '%v' ", result)
	}
}
