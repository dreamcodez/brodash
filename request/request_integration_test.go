package request_test

import (
	"strings"

	"github.com/dreamcodez/brodash/request"
	"github.com/dreamcodez/brodash/result"
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
			results := request.ParReq(reqs)
			slices.SortFunc(results, func(a, b result.Result[*resty.Response]) bool {
				return strings.Compare(a.Val.Request.URL, b.Val.Request.URL) == -1
			})
			Expect(results.Ok()).To(BeTrue())
			Expect(len(results)).To(Equal(3))
			vals := results.Values()
			Expect(vals[0].Body()).To(ContainSubstring(`"name":"Tatooine"`))
			Expect(vals[1].Body()).To(ContainSubstring(`"name":"Alderaan"`))
			Expect(vals[2].Body()).To(ContainSubstring(`"name":"Yavin IV"`))
		})
	})
})
