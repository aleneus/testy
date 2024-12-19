// Package testy implements some helpers for writing tests.
package testy

import (
	"testing"
)

// AssertPanic asserts that call of f gives panic.
func AssertPanic(t *testing.T, f func()) {
	t.Helper()

	havePanic := false
	defer func() {
		if r := recover(); r != nil {
			havePanic = true
		}
	}()

	f()

	Assert(t, havePanic)
}
