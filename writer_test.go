package testy

import (
	"testing"
)

func TestWriter(t *testing.T) {
	t.Parallel()

	w := new(Writer)

	w.Write([]byte("Hello"))
	AssertEqual(t, w.Pop(), "Hello")
	AssertEqual(t, w.Pop(), "")

	w.Write([]byte("world"))
	AssertEqual(t, w.Pop(), "world")
	AssertEqual(t, w.Pop(), "")
}
