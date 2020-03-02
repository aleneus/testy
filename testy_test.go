package testy

import (
	"testing"
)

func TestOk_Equal(t *testing.T) {
	AssertEqual(t, 1, 1)
	AssertEqual(t, 1, 1, "values differ")
}

func TestFail_Equal(t *testing.T) {
	AssertEqual(t, 1, 2, "values differ")
}
