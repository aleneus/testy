package testy

import "testing"

func TestAssertEqualComplex(t *testing.T) {
	t.Parallel()

	AssertEqualComplex(t, 1+2i, 1+2i, 0.0001)
}
