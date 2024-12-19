package testy

import "testing"

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
