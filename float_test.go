package testy

import "testing"

func TestAssertEqualFloat32(t *testing.T) {
	t.Parallel()

	AssertEqualFloat32(t, 1.12346, 1.12345, 0.0001)
}

func TestAssertEqualFloat64(t *testing.T) {
	t.Parallel()

	AssertEqualFloat64(t, 1.12346, 1.12345, 0.0001)
}

func TestAssertNotEqualFloat32(t *testing.T) {
	t.Parallel()

	AssertNotEqualFloat32(t, 1.2, 1.3, 0.00001)
}

func TestAssertNotEqualFloat64(t *testing.T) {
	t.Parallel()

	AssertNotEqualFloat64(t, 1.2, 1.3, 0.00001)
}
