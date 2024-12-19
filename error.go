// Package testy implements some helpers for writing tests.
package testy

import (
	"errors"
	"testing"
)

// AssertNoErr check if error is nil.
func AssertNoErr(t *testing.T, err error, msg ...interface{}) {
	t.Helper()

	if err != nil {
		t.Fatal(err, msg)
	}
}

// AssertError checks that error is not nil.
func AssertError(t *testing.T, err error, msg ...interface{}) {
	t.Helper()

	if err == nil {
		t.Fatal("No error", msg)
	}
}

// AssertError checks that error is wanted.
func AssertErrorIs(t *testing.T, err error, target error, msg ...interface{}) {
	t.Helper()

	if err == nil {
		t.Fatal("No error", msg)
	}

	if !errors.Is(err, target) {
		t.Fatal("Wrong error: ", err, msg)
	}
}
