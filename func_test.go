package testy

import (
	"errors"
	"testing"
)

func TestCall(t *testing.T) {
	t.Parallel()

	fn := func() (int, error) {
		return 1, nil
	}

	AssertEqual(t, Call(t, fn), 1)
}

func TestCallErr(t *testing.T) {
	t.Parallel()

	err := errors.New("error")

	t.Run("any error", func(t *testing.T) {
		t.Parallel()

		fn := func() (int, error) {
			return 1, err
		}

		CallErr(t, fn, nil)
	})

	t.Run("concrete error", func(t *testing.T) {
		t.Parallel()

		fn := func() (int, error) {
			return 1, err
		}

		CallErr(t, fn, err)
	})
}

func TestApply(t *testing.T) {
	t.Parallel()

	bs := []byte{1, 2, 3}
	fn := func(bs []byte) (int, error) {
		return len(bs), nil
	}

	AssertEqual(t, Apply(t, fn, bs), len(bs))
}

func TestApplyErr(t *testing.T) {
	t.Parallel()

	err := errors.New("error")

	t.Run("any error", func(t *testing.T) {
		t.Parallel()

		fn := func(bs []byte) (int, error) {
			return len(bs), err
		}

		ApplyErr(t, fn, []byte{1, 2, 3}, nil)
	})

	t.Run("concrete error", func(t *testing.T) {
		t.Parallel()

		fn := func(bs []byte) (int, error) {
			return len(bs), err
		}

		ApplyErr(t, fn, []byte{1, 2, 3}, err)
	})
}
