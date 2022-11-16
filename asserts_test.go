package testy

import (
	"fmt"
	"testing"
)

func TestAssertEqual(t *testing.T) {
	t.Parallel()

	AssertEqual(t, 1, 1)
	AssertEqual(t, "hello", "hello")
}

func TestAssertNotEqual(t *testing.T) {
	t.Parallel()

	AssertNotEqual(t, 1, 2)
	AssertNotEqual(t, "hello", "hallo")
}

func TestAssert(t *testing.T) {
	t.Parallel()

	Assert(t, 1 == 1)
}

func TestAssertTrue(t *testing.T) {
	t.Parallel()

	AssertTrue(t, 1 == 1)
}

func TestAssertFalse(t *testing.T) {
	t.Parallel()

	AssertFalse(t, 1 == 2)
}

func TestAssertError(t *testing.T) {
	t.Parallel()

	AssertError(t, fmt.Errorf("error"))
}

func TestAssertNoErr(t *testing.T) {
	t.Parallel()

	AssertNoErr(t, nil)
}

func TestAssertNotNil(t *testing.T) {
	t.Parallel()

	AssertNotNil(t, 1)
	AssertNotNil(t, "hello")
}

func TestAssertPanic(t *testing.T) {
	t.Parallel()

	f := func() int {
		a := 1
		return 1 / (1 - a)
	}

	AssertPanic(t, func() { f() })
}
