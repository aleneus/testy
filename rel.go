// Package testy implements some helpers for writing tests.
package testy

// Based on: https://floating-point-gui.de/errors/comparison/

import (
	"math"
	"testing"
)

// AssertRelEqualFloat32 checks that two float32 values are equal with relative eps.
func AssertRelEqualFloat32(t *testing.T, v1, v2 float32, eps float64, msg ...interface{}) {
	t.Helper()

	if !nearlyEqual(float64(v1), float64(v2), eps, float64(minNormal32), float64(math.MaxFloat32)) {
		t.Fatal(v1, "!=", v2, msg)
	}
}

// AssertRelNotEqualFloat32 checks that two float32 values are not equal with relative eps.
func AssertRelNotEqualFloat32(t *testing.T, v1, v2 float32, eps float64, msg ...interface{}) {
	t.Helper()

	if nearlyEqual(float64(v1), float64(v2), eps, float64(minNormal32), float64(math.MaxFloat32)) {
		t.Fatal(v1, "!=", v2, msg)
	}
}

const minNormal32 = float32(0x1p-126)

// AssertRelEqualFloat64 checks that two float64 values are equal with relative eps.
func AssertRelEqualFloat64(t *testing.T, v1, v2 float64, eps float64, msg ...interface{}) {
	t.Helper()

	if !nearlyEqual(v1, v2, eps, minNormal64, math.MaxFloat64) {
		t.Fatal(v1, "!=", v2, msg)
	}
}

// AssertRelNotEqualFloat64 checks that two float64 values are not equal.
func AssertRelNotEqualFloat64(t *testing.T, v1, v2 float64, eps float64, msg ...interface{}) {
	t.Helper()

	if nearlyEqual(v1, v2, eps, minNormal64, math.MaxFloat64) {
		t.Fatal(v1, "!=", v2, msg)
	}
}

const minNormal64 = 0x1p-1022

func nearlyEqual(a, b, eps, minNormal, maxFloat float64) bool {
	absA := math.Abs(float64(a))
	absB := math.Abs(float64(b))

	if a == b {
		return true
	}

	diff := math.Abs(float64(a) - float64(b))

	if a == 0 || b == 0 || (absA+absB < minNormal) {
		return diff < (eps * minNormal)
	}

	return diff/math.Min((absA+absB), maxFloat) < eps
}
