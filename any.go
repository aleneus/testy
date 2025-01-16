// Package testy implements some helpers for writing tests.
package testy

import (
	"reflect"
	"testing"
)

// AssertEqual checks that two values are equal.
func AssertEqual(t *testing.T, v1, v2 any, msg ...any) {
	t.Helper()

	if v1 != v2 {
		t.Fatal(v1, "!=", v2, msg)
	}
}

// AssertNotEqual checks that two values are not equal.
func AssertNotEqual(t *testing.T, v1, v2 any, msg ...any) {
	t.Helper()

	if v1 == v2 {
		t.Fatal("Both values are equal to", v1, msg)
	}
}

// AssertNotNil checks that value is not nil.
func AssertNotNil(t *testing.T, v any, msg ...any) {
	t.Helper()

	if v == nil {
		t.Fatal("is nil", msg)
	}

	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		return
	}

	if reflect.ValueOf(v).IsNil() {
		t.Fatal("is nil", msg)
	}
}
