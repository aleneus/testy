// Package testy implements some helpers for writing tests
package testy

import (
	"testing"
)

// AssertEqual checks that two values are equal
func AssertEqual(t *testing.T, v1, v2 interface{}) {
	if v1 != v2 {
		t.Fatal(v1, "!=", v2)
	}
}

// AssertNotEqual checks that two values are not equal
func AssertNotEqual(t *testing.T, v1, v2 interface{}) {
	if v1 == v2 {
		t.Fatal(v1, "==", v2)
	}
}

// AssertTrue checks that value is true
func AssertTrue(t *testing.T, v bool) {
	if !v {
		t.Fatal("false is not true")
	}
}

// AssertFalse checks that value is false
func AssertFalse(t *testing.T, v bool) {
	if v {
		t.Fatal("true is not false")
	}
}

// AssertNoErr check if error is not nil
func AssertNoErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

// AssertError checks that error is nil
func AssertError(t *testing.T, err error) {
	if err == nil {
		t.Fatal("No error")
	}
}
