package nrfasthttp

import (
	"time"

	fasthttp "gopkg.in/valyala/fasthttp.v20160316"
)

type Client interface {
	Do(req *fasthttp.Request, res *fasthttp.Response) error
	DoTimeout(req *fasthttp.Request, res *fasthttp.Response, timeout time.Duration) error
	DoDeadline(req *fasthttp.Request, res *fasthttp.Response, deadline time.Time) error
}
