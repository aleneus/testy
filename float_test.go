package testy

import (
	"math"
	"testing"
)

func TestAssertEqualFloat32(t *testing.T) {
	t.Parallel()

	AssertEqualFloat32(t, 1.12346, 1.12345, 0.0001)
	AssertEqualFloat32(t, 0.000000001000001, 0.000000001000002, 0.00001)
	AssertEqualFloat32(t, 0.000000001000002, 0.000000001000001, 0.00001)
	AssertEqualFloat32(t, 0.0, 0.0, 0.00001)
	AssertEqualFloat32(t, 0.0, 1e-40, 0.1)
	AssertEqualFloat32(t, 1e-40, 0.0, 0.1)
	AssertEqualFloat32(t, 0.0, -1e-40, 0.1)
	AssertEqualFloat32(t, -1e-40, 0.0, 0.1)
	AssertEqualFloat32(t, 0.0, -1e-40, 0.000001)
	AssertEqualFloat32(t, -1e-40, 0.0, 0.000001)
}

func TestAssertNotEqualFloat32(t *testing.T) {
	t.Parallel()

	AssertNotEqualFloat32(t, 1.2, 1.3, 0.00001)
	AssertNotEqualFloat32(t, float32(math.NaN()), 1, 0.00001)
	AssertNotEqualFloat32(t, float32(math.NaN()), float32(math.NaN()), 0.00001)
}

func TestAssertEqualFloat64(t *testing.T) {
	t.Parallel()

	AssertEqualFloat64(t, 1.12346, 1.12345, 0.0001)
	AssertEqualFloat64(t, 0.0000000010000001, 0.0000000010000002, 0.0000001)
	AssertEqualFloat64(t, 0.0000000010000002, 0.0000000010000001, 0.0000001)
	AssertEqualFloat64(t, 0.0, 0.0, 0.00001)
	AssertEqualFloat64(t, 0.0, 1e-309, 0.1)
	AssertEqualFloat64(t, 1e-309, 0.0, 0.1)
	AssertEqualFloat64(t, 0.0, -1e-309, 0.1)
	AssertEqualFloat64(t, -1e-309, 0.0, 0.1)
	AssertEqualFloat64(t, 0.0, -1e-309, 0.000001)
	AssertEqualFloat64(t, -1e-309, 0.0, 0.000001)
}

func TestAssertNotEqualFloat64(t *testing.T) {
	t.Parallel()

	AssertNotEqualFloat64(t, 1.2, 1.3, 0.00001)
	AssertNotEqualFloat64(t, math.NaN(), 1, 0.00001)
	AssertNotEqualFloat64(t, math.NaN(), math.NaN(), 0.00001)
}
