package lists_test

import (
	"strconv"

	"github.com/dreamcodez/brodash/lists"
	"github.com/dreamcodez/brodash/result"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/exp/slices"
)

var _ = Describe("lists", func() {
	Context("Map(...)", func() {
		shouldMapPlusOneAdder(lists.Map[int, int])()
	})

	Context("ParMap(...)", func() {
		shouldMapPlusOneAdder(lists.ParMap[int, int])()
	})

	Context("MapWithError(...)", func() {
		shouldMapAtoi(lists.MapWithError[string, int])()
	})

	PContext("ParMapWithError(...)", func() {
	})

	Context("ParMapWithResults(...)", func() {
		shouldMapAtoiResults(lists.ParMapWithResults[string, int])()

	})
})

func shouldMapPlusOneAdder(fn func(lst []int, fn lists.MapFn[int, int]) []int) func() {
	return func() {
		It("should be able to map a function which adds one to each number over a list of integers", func() {
			lst := []int{1, 2, 3}
			expected := []int{2, 3, 4}

			actual := fn(lst, func(i int) int {
				return i + 1
			})
			slices.Sort(actual)

			Expect(actual).To(Equal(expected))
		})
	}
}

func shouldMapAtoi(fn func(lst []string, fn lists.MapErrFn[string, int]) ([]int, error)) func() {
	return func() {
		It("should be able to map Atoi over a list of string and produce a list of integers", func() {
			lst := []string{"1", "2", "3"}
			expected := []int{1, 2, 3}

			actual, err := fn(lst, strconv.Atoi)
			Expect(err).To(BeNil())
			slices.Sort(actual)

			Expect(actual).To(Equal(expected))
		})
	}
}

func shouldMapAtoiResults(fn func(lst []string, fn lists.MapErrFn[string, int]) result.Results[int]) func() {
	return func() {
		shouldMapAtoi(func(lst []string, fn lists.MapErrFn[string, int]) ([]int, error) {
			return lists.ParMapWithResults(lst, fn).Get()
		})()
	}
}
