// Package testy provides the assert helpers for writing unit tests.
package testy

import (
	"math"
	"testing"
)

// AssertEqualFloat32 checks that two float32 values are equal with absolute eps.
func AssertEqualFloat32(t *testing.T, v1, v2 float32, eps float64, msg ...any) {
	t.Helper()

	if math.Abs(float64(v1)-float64(v2)) > eps {
		t.Fatal(v1, "!=", v2, msg)
	}
}

// AssertNotEqualFloat32 checks that two float32 values are not equal with absolute eps.
func AssertNotEqualFloat32(t *testing.T, v1, v2 float32, eps float64, msg ...any) {
	t.Helper()

	if math.Abs(float64(v1)-float64(v2)) <= eps {
		t.Fatal(v1, "==", v2, msg)
	}
}

// AssertEqualFloat64 checks that two float64 values are equal with absolute eps.
func AssertEqualFloat64(t *testing.T, v1, v2 float64, eps float64, msg ...any) {
	t.Helper()

	if math.Abs(v1-v2) > eps {
		t.Fatal(v1, "!=", v2, msg)
	}
}

// AssertNotEqualFloat64 checks that two float64 values are not equal with absolute eps.
func AssertNotEqualFloat64(t *testing.T, v1, v2 float64, eps float64, msg ...any) {
	t.Helper()

	if math.Abs(v1-v2) <= eps {
		t.Fatal(v1, "==", v2, msg)
	}
}
