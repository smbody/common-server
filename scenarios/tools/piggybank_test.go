package tools

import (
	"testing"
	"github.com/smbody/common-server/application"
)

func TestCalculate(t *testing.T) {
	t.Log("Piggybank just sum two values.")
	p := NewPiggyBank()
	if result := p.Calculate(application.IntToDecimal(150), application.IntToDecimal(50)); result.Cmp(application.IntToDecimal(200)) != 0 {
		t.Errorf("Expected result: '%v' but was '%v'. ", 200, result)
	}
}
