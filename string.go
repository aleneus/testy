// Package testy implements some helpers for writing tests.
package testy

import (
	"strings"
	"testing"
)

// AssertSubstr asserts that first string is substring of the
// second one.
func AssertSubstr(t *testing.T, substr, str string) {
	t.Helper()

	if !strings.Contains(str, substr) {
		t.Fatal("there is no", substr, "in", str)
	}
}

// AssertNotSubstr asserts that first string is not substring of the
// second one.
func AssertNotSubstr(t *testing.T, substr, str string) {
	t.Helper()

	if strings.Contains(str, substr) {
		t.Fatal("there is", substr, "in", str)
	}
}
