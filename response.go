package nrfasthttp

import (
	"github.com/newrelic/go-agent/http"
	fasthttp "gopkg.in/valyala/fasthttp.v20160316"
)

type Response struct {
	req *fasthttp.Request
	res *fasthttp.Response
}

func NewResponse(res *fasthttp.Response, req *fasthttp.Request) Response {
	return Response{
		req: req,
		res: res,
	}
}

func (r Response) Header() http.Header {
	return ResponseHeader{&r.res.Header}
}

func (r Response) Code() int {
	return r.res.StatusCode()
}

func (r Response) Request() http.Request {
	return Request{r.req}
}
