package testy

import (
	"testing"
)

func TestAssertPanic(t *testing.T) {
	t.Parallel()

	f := func() int {
		a := 1
		return 1 / (1 - a)
	}

	AssertPanic(t, func() { f() })
}
