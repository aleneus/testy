package testy

import "testing"

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

func TestAssertNotNil(t *testing.T) {
	t.Parallel()

	AssertNotNil(t, 1)
	AssertNotNil(t, "hello")
}
