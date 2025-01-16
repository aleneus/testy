// Package testy implements some helpers for writing tests.
package testy

import "testing"

// Assert checks that value is true.
func Assert(t *testing.T, v bool, msg ...any) {
	t.Helper()

	AssertTrue(t, v, msg)
}

// AssertTrue checks that value is true.
func AssertTrue(t *testing.T, v bool, msg ...any) {
	t.Helper()

	if !v {
		t.Fatal("false is not true", msg)
	}
}

// AssertFalse checks that value is false.
func AssertFalse(t *testing.T, v bool, msg ...any) {
	t.Helper()

	if v {
		t.Fatal("true is not false", msg)
	}
}
