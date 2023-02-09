package result

import "github.com/hashicorp/go-multierror"

type Results[T any] []Result[T]

func (rs Results[T]) Get() ([]T, error) {
	err := rs.MultiError()
	if err != nil {
		return nil, err
	}
	return rs.Values(), nil
}

func (rs Results[T]) Ok() bool {
	return rs.MultiError() == nil
}

func (rs Results[T]) Values() []T {
	var vals []T
	for _, r := range rs {
		vals = append(vals, r.Val)
	}
	return vals
}

func (rs Results[T]) Errors() []error {
	var errs []error
	for _, r := range rs {
		errs = append(errs, r.Err)
	}
	return errs
}

func (rs Results[T]) MultiError() error {
	var multierr error
	for _, err := range rs.Errors() {
		if err != nil {
			multierr = multierror.Append(multierr, err)
		}
	}
	return multierr
}

type Result[T any] struct {
	Val T
	Err error
}

func NewResult[T any](val T, err error) Result[T] {
	return Result[T]{
		Val: val,
		Err: err,
	}
}

func (r Result[T]) Get() (T, error) {
	return r.Val, r.Err
}

func (r Result[T]) Ok() bool {
	return r.Err == nil
}

func (r Result[T]) ValueOr(v T) T {
	if r.Ok() {
		return r.Val
	}
	return v
}
