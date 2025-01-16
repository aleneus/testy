package testy

import "testing"

func TestAssertRelNotEqualFloat32(t *testing.T) {
	t.Parallel()

	AssertRelNotEqualFloat32(t, 0.000000000001001, 0.000000000001002, 0.00001)
	AssertRelNotEqualFloat32(t, 0.000000000001002, 0.000000000001001, 0.00001)
	AssertRelNotEqualFloat32(t, 0.0, 1e-40, 0.000001)
	AssertRelNotEqualFloat32(t, 1e-40, 0.0, 0.000001)
}

func TestAssertRelNotEqualFloat64(t *testing.T) {
	t.Parallel()

	AssertRelNotEqualFloat64(t, 0.000000000001001, 0.000000000001002, 0.0000001)
	AssertRelNotEqualFloat64(t, 0.000000000001002, 0.000000000001001, 0.0000001)
	AssertRelNotEqualFloat64(t, 0.0, 1e-309, 0.000001)
	AssertRelNotEqualFloat64(t, 1e-309, 0.0, 0.000001)
}
