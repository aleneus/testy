// Package testy provides the assert helpers for writing unit tests.
package testy

import "testing"

// Call invokes fn, fails the test if it returns an error, and returns
// its result.
func Call[Y any](t *testing.T, fn func() (Y, error)) Y {
	t.Helper()

	res, err := fn()
	AssertNoErr(t, err)

	return res
}

// CallErr invokes fn, fails the test if it returns no error,
// optionally checks the error against err, and returns it.
func CallErr[Y any](t *testing.T, fn func() (Y, error), err error) error {
	t.Helper()

	_, e := fn()
	AssertError(t, e)

	if err != nil {
		AssertErrorIs(t, e, err)
	}

	return e
}

// Apply invokes fn with x, fails the test if it returns an error, and
// returns its result.
func Apply[X, Y any](t *testing.T, fn func(x Y) (X, error), x Y) X {
	t.Helper()

	return Call(t, func() (X, error) {
		return fn(x)
	})
}

// ApplyErr invokes fn with x, fails the test if it returns no error,
// optionally checks the error against err, and returns it.
func ApplyErr[T, F any](t *testing.T, fn func(x F) (T, error), x F, err error) error {
	t.Helper()

	return CallErr(t, func() (T, error) {
		return fn(x)
	}, err)
}
