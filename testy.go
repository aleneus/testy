// Package testy implements some helpers for writing tests
package testy

import (
	"testing"
)

// AssertEqual checks that two values are equal
func AssertEqual(t *testing.T, v1, v2 interface{}, msg ...interface{}) {
	if v1 != v2 {
		t.Fatal(v1, "!=", v2, msg)
	}
}

// AssertNotEqual checks that two values are not equal
func AssertNotEqual(t *testing.T, v1, v2 interface{}, msg ...interface{}) {
	if v1 == v2 {
		t.Fatal(v1, "==", v2, msg)
	}
}

// AssertTrue checks that value is true
func AssertTrue(t *testing.T, v bool, msg ...interface{}) {
	if !v {
		t.Fatal("false is not true", msg)
	}
}

// AssertFalse checks that value is false
func AssertFalse(t *testing.T, v bool, msg ...interface{}) {
	if v {
		t.Fatal("true is not false", msg)
	}
}

// AssertNoErr check if error is not nil
func AssertNoErr(t *testing.T, err error, msg ...interface{}) {
	if err != nil {
		t.Fatal(err, msg)
	}
}

// AssertError checks that error is nil
func AssertError(t *testing.T, err error, msg ...interface{}) {
	if err == nil {
		t.Fatal("No error", msg)
	}
}
