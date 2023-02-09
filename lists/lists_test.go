package lists_test

import (
	"strconv"

	"github.com/dreamcodez/brodash/lists"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/exp/slices"
)

var _ = Describe("lists", func() {
	Context("Map(...)", func() {
		It("should be able to map a function which adds one to each number over a list of integers", func() {
			lst := []int{1, 2, 3}
			expected := []int{2, 3, 4}

			actual := lists.Map(lst, func(i int) int {
				return i + 1
			})
			Expect(actual).To(Equal(expected))
		})
	})

	Context("ParMap(...)", func() {
		It("should be able to map a function which adds one to each number over a list of integers", func() {
			lst := []int{1, 2, 3}
			expected := []int{2, 3, 4}

			actual := lists.Map(lst, func(i int) int {
				return i + 1
			})
			Expect(actual).To(Equal(expected))
		})
	})

	Context("MapWithError(...)", func() {
		It("should be able to map Atoi over a list of string and produce a list of integers", func() {
			lst := []string{"1", "2", "3"}
			expected := []int{1, 2, 3}

			actual, err := lists.MapWithError(lst, strconv.Atoi)
			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})
	})

	Context("ParMapWithResults(...)", func() {
		It("should be able to map Atoi over a list of string and produce a list of integers", func() {
			lst := []string{"1", "2", "3"}
			expected := []int{1, 2, 3}
			actual, err := lists.ParMapWithResults(lst, strconv.Atoi).Get()
			Expect(err).To(BeNil())
			slices.Sort(actual)
			Expect(actual).To(Equal(expected))
		})
	})
})
