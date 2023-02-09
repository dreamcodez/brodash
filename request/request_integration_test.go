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
				mkReq("GET", "https://swapi.dev/api/planets/1"),
				mkReq("GET", "https://swapi.dev/api/planets/2"),
				mkReq("GET", "https://swapi.dev/api/planets/3"),
			}
			results := request.ParReq(reqs)
			slices.SortFunc(results, func(a, b result.Result[*resty.Response]) bool {
				return strings.Compare(a.Val.Request.URL, b.Val.Request.URL) == -1
			})
			Expect(results.Ok()).To(BeTrue())
			Expect(len(results)).To(Equal(3))
			Expect(results[0].Val.Body()).To(ContainSubstring(`"name":"Tatooine"`))
			Expect(results[1].Val.Body()).To(ContainSubstring(`"name":"Alderaan"`))
			Expect(results[2].Val.Body()).To(ContainSubstring(`"name":"Yavin IV"`))
		})
	})
})

func mkReq(method string, url string) *resty.Request {
	rc := request.Client
	req := rc.NewRequest()
	req.Method = "GET"
	req.URL = url
	return req
}
