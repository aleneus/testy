package testy

import (
	"fmt"
	"testing"
)

func TestAssertEqual(t *testing.T) {
	AssertEqual(t, 1, 1)
	AssertEqual(t, "hello", "hello")
}

func TestAssertNotEqual(t *testing.T) {
	AssertNotEqual(t, 1, 2)
	AssertNotEqual(t, "hello", "hallo")
}

func TestAssert(t *testing.T) {
	Assert(t, 1 == 1)
}

func TestAssertTrue(t *testing.T) {
	AssertTrue(t, 1 == 1)
}

func TestAssertFalse(t *testing.T) {
	AssertFalse(t, 1 == 2)
}

func TestAssertError(t *testing.T) {
	AssertError(t, fmt.Errorf("error"))
}

func TestAssertNoErr(t *testing.T) {
	AssertNoErr(t, nil)
}

func TestAssertNotNil(t *testing.T) {
	AssertNotNil(t, 1)
	AssertNotNil(t, "hello")
}
