package testy

import (
	"errors"
	"fmt"
	"testing"
)

func TestAssertNoErr(t *testing.T) {
	t.Parallel()

	AssertNoErr(t, nil)
}

func TestAssertError(t *testing.T) {
	t.Parallel()

	AssertError(t, fmt.Errorf("error"))
}

func TestAssertErrorIs(t *testing.T) {
	t.Parallel()

	errTarget := errors.New("target")
	AssertErrorIs(t, errTarget, errTarget)
}
