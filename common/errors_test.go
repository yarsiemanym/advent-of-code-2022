package common

import (
	"errors"
	"testing"
)

func Test_check_Error(t *testing.T) {
	err := errors.New("This error is expected.")
	assertPanic(t, func() { Check(err) })
}

func Test_check_NoError(t *testing.T) {
	Check(nil)
}

func assertPanic(t *testing.T, fn func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic.")
		}
	}()

	fn()
}
