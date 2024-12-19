// Package testy implements some helpers for writing tests.
package testy

import "testing"

// AssertEqualComplex checks that two complex values are almost equal.
func AssertEqualComplex(t *testing.T, c1, c2 complex128, eps float64, msg ...interface{}) {
	t.Helper()

	AssertEqualFloat64(t, real(c1), real(c2), eps, msg)
	AssertEqualFloat64(t, imag(c1), imag(c2), eps, msg)
}
