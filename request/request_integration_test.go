package request_test

import (
	"strings"

	"github.com/dreamcodez/brodash/request"
	"github.com/dreamcodez/brodash/result"
	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Request Integration", func() {
	Context("ParReq(...)", func() {
		Context("while requesting 3 planets on swapi", func() {
			reqs := []*resty.Request{
				request.R("GET", "https://swapi.dev/api/planets/1"),
				request.R("GET", "https://swapi.dev/api/planets/2"),
				request.R("GET", "https://swapi.dev/api/planets/3"),
			}

			results := request.ParReq[map[string]interface{}](reqs)
			results.Sort(func(a, b result.Result[map[string]interface{}]) bool {
				return strings.Compare(a.Val["name"].(string), b.Val["name"].(string)) == -1
			})
			responses, err := results.Get()

			It("the response should be successful", func() {
				Expect(err).To(BeNil())
			})

			It("the response should have a proper length", func() {
				Expect(len(responses)).To(Equal(3))
			})

			It("the marshalled json response should contain the correct name values", func() {
				Expect(responses[0]["name"]).To(ContainSubstring(`Alderaan`))
				Expect(responses[1]["name"]).To(ContainSubstring(`Tatooine`))
				Expect(responses[2]["name"]).To(ContainSubstring(`Yavin IV`))
			})
		})
	})
	Context("ParRawReq(...)", func() {
		Context("while requesting 3 planets on swapi", func() {
			reqs := []*resty.Request{
				request.R("GET", "https://swapi.dev/api/planets/1"),
				request.R("GET", "https://swapi.dev/api/planets/2"),
				request.R("GET", "https://swapi.dev/api/planets/3"),
			}

			results := request.ParReqRaw(reqs)
			results.Sort(func(a, b result.Result[*resty.Response]) bool {
				return strings.Compare(a.Val.Request.URL, b.Val.Request.URL) == -1
			})

			responses, err := results.Get()

			It("the response should be successful", func() {
				Expect(err).To(BeNil())
			})

			It("the response should have a proper length", func() {
				Expect(len(responses)).To(Equal(3))
			})

			It("the response should contain name segments", func() {
				Expect(len(responses)).To(Equal(3))
				Expect(responses[0].Body()).To(ContainSubstring(`"name":"Tatooine"`))
				Expect(responses[1].Body()).To(ContainSubstring(`"name":"Alderaan"`))
				Expect(responses[2].Body()).To(ContainSubstring(`"name":"Yavin IV"`))
			})
		})
	})
})
