package mocks

import (
	"testing"
	"github.com/smbody/common-server/app"
)

func AwaitPanic(t *testing.T, fn func(), errs ...app.ErrorCode) {

	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok && len(errs) > 0 {
				checkError(t, app.GetError(errs[0]), err)
			}
		} else {
			t.Error("The code did not panic")
		}
	}()

	fn()
}
func checkError(t *testing.T, expected error, existed error) {
	if expected != existed {
		t.Errorf("Expected %q but has %q", expected, existed)
	}
}


