package calculator

import (
"testing"
	"github.com/smbody/common-server/application"
	"github.com/smbody/common-server/scenarios/tools"
)

func TestAddIteration(t *testing.T) {
	t.Log("After iteration value must grow.")
	invest := application.IntToDecimal(50)
	p := tools.NewPiggyBank()
	d := &tools.Distribution{p, application.IntToDecimal(100), invest}
	i := NewToolIterator(d)
	for c :=1; c<=5; c++ {
		if amount := i.AddIteration(); i.value.Cmp(invest.Mul(application.IntToDecimal(c))) != 0 {
			t.Errorf("Iteration %v(amount=%v) expected result: '%v' but was '%v'.", c, amount, invest.Mul(application.IntToDecimal(c)), i.value,)
		}
	}
}

