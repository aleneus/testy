// Package testy implements some helpers for writing tests
package testy

import (
	"fmt"
	"testing"
)

// AssertTrue checks that value is true
func AssertTrue(t *testing.T, v bool, msg ...interface{}) {
	if !v {
		fmt.Println(info())
		t.Fatal("false is not true", msg)
	}
}

// AssertFalse checks that value is false
func AssertFalse(t *testing.T, v bool, msg ...interface{}) {
	if v {
		fmt.Println(info())
		t.Fatal("true is not false", msg)
	}
}

// AssertEqual checks that two values are equal
func AssertEqual(t *testing.T, v1, v2 interface{}, msg ...interface{}) {
	if v1 != v2 {
		fmt.Println(info())
		t.Fatal(v1, "!=", v2, msg)
	}
}

// AssertNotEqual checks that two values are not equal
func AssertNotEqual(t *testing.T, v1, v2 interface{}, msg ...interface{}) {
	if v1 == v2 {
		fmt.Println(info())
		t.Fatal("Both values are equal to", v1, msg)
	}
}

// AssertNoErr check if error is not nil
func AssertNoErr(t *testing.T, err error, msg ...interface{}) {
	if err != nil {
		fmt.Println(info())
		t.Fatal(err, msg)
	}
}

// AssertError checks that error is nil
func AssertError(t *testing.T, err error, msg ...interface{}) {
	if err == nil {
		fmt.Println(info())
		t.Fatal("No error", msg)
	}
}
