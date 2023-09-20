// Package testy implements some helpers for writing tests.
package testy

import (
	"math"
	"testing"
)

// AssertEqualFloat32 checks that two float32 values are equal.
func AssertEqualFloat32(t *testing.T, v1, v2 float32, eps float64, msg ...interface{}) {
	t.Helper()

	f1 := float64(v1)
	f2 := float64(v2)

	if math.IsNaN(f1) || math.IsNaN(f2) {
		t.Fatal(v1, "!=", v2, msg)
	}

	if math.Abs(f1-f2) >= eps {
		t.Fatal(v1, "!=", v2, msg)
	}
}

// AssertEqualFloat64 checks that two float64 values are equal.
func AssertEqualFloat64(t *testing.T, v1, v2 float64, eps float64, msg ...interface{}) {
	t.Helper()

	if math.IsNaN(v1) || math.IsNaN(v2) {
		t.Fatal(v1, "!=", v2, msg)
	}

	if math.Abs(v1-v2) >= eps {
		t.Fatal(v1, "!=", v2, msg)
	}
}

// AssertNotEqualFloat32 checks that two float32 values are not equal.
func AssertNotEqualFloat32(t *testing.T, v1, v2 float32, eps float64, msg ...interface{}) {
	t.Helper()

	if math.Abs(float64(v1)-float64(v2)) < eps {
		t.Fatal(v1, "==", v2, msg)
	}
}

// AssertNotEqualFloat64 checks that two float64 values are not equal.
func AssertNotEqualFloat64(t *testing.T, v1, v2 float64, eps float64, msg ...interface{}) {
	t.Helper()

	if math.Abs(v1-v2) < eps {
		t.Fatal(v1, "==", v2, msg)
	}
}
