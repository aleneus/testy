package testy

import (
	"testing"
)

func TestFail_LogAssertEqual(t *testing.T) {
	if LogAssertEqual(1, 2, "values differ") {
		t.Fatal("Worn result")
	}
}
