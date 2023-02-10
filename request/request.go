package request

import (
	"github.com/dreamcodez/brodash/lists"
	"github.com/dreamcodez/brodash/result"
	"github.com/go-resty/resty/v2"
)

var Client = resty.New()

func R(method string, url string) *resty.Request {
	req := Client.R()
	req.Method = "GET"
	req.URL = url
	return req
}

func ParReq[Response any](requests []*resty.Request) result.Results[Response] {
	return ParReqWithRestyClient[Response](Client, requests)
}

func ParReqWithRestyClient[Out any](rc *resty.Client, requests []*resty.Request) result.Results[Out] {
	return lists.ParMapWithResults(requests, func(request *resty.Request) (Out, error) {
		var out Out

		request.SetResult(out)
		resp, err := request.Send()
		if err != nil {
			return out, err
		}

		out = *resp.Result().(*Out)
		return out, nil
	})
}

func ParReqRaw(requests []*resty.Request) result.Results[*resty.Response] {
	return ParReqRawWithRestyClient(Client, requests)
}

func ParReqRawWithRestyClient(rc *resty.Client, requests []*resty.Request) result.Results[*resty.Response] {
	return lists.ParMapWithResults(requests, func(request *resty.Request) (*resty.Response, error) {
		return request.Send()
	})
}
