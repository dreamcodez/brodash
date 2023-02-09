package lists

import (
	"github.com/dreamcodez/brodash/result"
)

type MapFn[In any, Out any] func(In) Out
type MapErrFn[In any, Out any] func(In) (Out, error)

// Map maps the fn to the array
func Map[In any, Out any](lst []In, fn MapFn[In, Out]) []Out {
	rval, _ := MapWithError(lst, func(val In) (Out, error) {
		return fn(val), nil
	})
	return rval
}

// MapWithResults maps the fn to the array and will return the results after all items have run
func MapWithResults[In any, Out any](lst []In, fn MapErrFn[In, Out]) result.Results[Out] {
	var out result.Results[Out]
	for _, val := range lst {
		out = append(out, result.NewResult(fn(val)))
	}
	return out
}

// MapWithError maps the fn to the array and will return error on first failure
func MapWithError[In any, Out any](lst []In, fn MapErrFn[In, Out]) ([]Out, error) {
	var out []Out
	for _, val := range lst {
		rval, err := fn(val)
		if err != nil {
			return nil, err
		}
		out = append(out, rval)
	}
	return out, nil
}

// ParMap maps the fn to the array in parallel -- ordering is not guaranteed
func ParMap[In any, Out any](lst []In, fn MapFn[In, Out]) []Out {
	results := ParMapWithResults(lst, func(val In) (Out, error) {
		return fn(val), nil
	})
	return results.Values()
}

// ParMapWithResults maps the fn to the array in parallel and will return a list of results -- ordering is not guaranteed
func ParMapWithResults[In any, Out any](lst []In, fn MapErrFn[In, Out]) result.Results[Out] {
	queue := make(chan result.Result[Out])
	var out result.Results[Out]
	for _, val := range lst {
		go func(val In) {
			queue <- result.NewResult(fn(val))
		}(val)
	}
	for i := 0; i < len(lst); i++ {
		out = append(out, <-queue)
	}
	return out
}
