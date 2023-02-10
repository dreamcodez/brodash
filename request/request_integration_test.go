package request_test

import (
	"strings"

	"github.com/dreamcodez/brodash/request"
	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/exp/slices"
)

var _ = Describe("Request Integration", func() {
	Context("ParReq(...)", func() {
		It("should request 3 planets on swapi successfully", func() {
			reqs := []*resty.Request{
				request.R("GET", "https://swapi.dev/api/planets/1"),
				request.R("GET", "https://swapi.dev/api/planets/2"),
				request.R("GET", "https://swapi.dev/api/planets/3"),
			}

			results, err := request.ParReq(reqs).Get()
			Expect(err).To(BeNil())
			slices.SortFunc(results, func(a, b *resty.Response) bool {
				return strings.Compare(a.Request.URL, b.Request.URL) == -1
			})

			Expect(len(results)).To(Equal(3))
			Expect(results[0].Body()).To(ContainSubstring(`"name":"Tatooine"`))
			Expect(results[1].Body()).To(ContainSubstring(`"name":"Alderaan"`))
			Expect(results[2].Body()).To(ContainSubstring(`"name":"Yavin IV"`))
		})
	})
})
