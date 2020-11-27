package testy

import (
	"testing"
)

func Test_info(t *testing.T) {
	if len(info()) == 0 {
		t.Fatal("info() gives nothing")
	}
}
