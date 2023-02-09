package lists

import (
	"github.com/dreamcodez/brodash/result"
)

// Map maps the fn to the array
func Map[In any, Out any](lst []In, fn func(val In) Out) []Out {
	rval, _ := MapWithError(lst, func(val In) (Out, error) {
		return fn(val), nil
	})
	return rval
}

// MapWithError maps the fn to the array and will return error on first failure
func MapWithError[In any, Out any](lst []In, fn func(val In) (Out, error)) ([]Out, error) {
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
func ParMap[In any, Out any](lst []In, fn func(In) Out) []Out {
	results := ParMapWithResults(lst, func(val In) (Out, error) {
		return fn(val), nil
	})
	return results.Values()
}

// ParMapWithResults maps the fn to the array in parallel and will return a list of results -- ordering is not guaranteed
func ParMapWithResults[In any, Out any](lst []In, fn func(In) (Out, error)) result.Results[Out] {
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
