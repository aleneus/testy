package testy

import "testing"

func TestAssertSubstr(t *testing.T) {
	t.Parallel()

	AssertSubstr(t, "aa", "bbaabb")
}

func TestAssertNotSubstr(t *testing.T) {
	t.Parallel()

	AssertNotSubstr(t, "aaa", "bbaabb")
}
