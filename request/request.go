package request

import (
	"github.com/dreamcodez/brodash/lists"
	"github.com/dreamcodez/brodash/result"
	"github.com/go-resty/resty/v2"
)

var Client = resty.New()

func ParReq(requests []*resty.Request) result.Results[*resty.Response] {
	return ParReqWithRestyClient(Client, requests)
}

func ParReqWithRestyClient(rc *resty.Client, requests []*resty.Request) result.Results[*resty.Response] {
	return lists.ParMapWithResults(requests, func(request *resty.Request) (*resty.Response, error) {
		return request.Send()
	})
}