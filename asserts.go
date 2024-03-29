// Package testy implements some helpers for writing tests.
package testy

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

// Assert checks that value is true.
func Assert(t *testing.T, v bool, msg ...interface{}) {
	t.Helper()

	AssertTrue(t, v, msg)
}

// AssertTrue checks that value is true.
func AssertTrue(t *testing.T, v bool, msg ...interface{}) {
	t.Helper()

	if !v {
		t.Fatal("false is not true", msg)
	}
}

// AssertFalse checks that value is false.
func AssertFalse(t *testing.T, v bool, msg ...interface{}) {
	t.Helper()

	if v {
		t.Fatal("true is not false", msg)
	}
}

// AssertEqual checks that two values are equal.
func AssertEqual(t *testing.T, v1, v2 interface{}, msg ...interface{}) {
	t.Helper()

	if v1 != v2 {
		t.Fatal(v1, "!=", v2, msg)
	}
}

// AssertNotEqual checks that two values are not equal.
func AssertNotEqual(t *testing.T, v1, v2 interface{}, msg ...interface{}) {
	t.Helper()

	if v1 == v2 {
		t.Fatal("Both values are equal to", v1, msg)
	}
}

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

// AssertNotNil checks that value is not nil.
func AssertNotNil(t *testing.T, v interface{}, msg ...interface{}) {
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
