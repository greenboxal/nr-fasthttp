package nrfasthttp

import (
	"time"

	"github.com/valyala/fasthttp"
)

type Client interface {
	Do(req *fasthttp.Request, res *fasthttp.Response) error
	DoTimeout(req *fasthttp.Request, res *fasthttp.Response, timeout time.Duration) error
	DoDeadline(req *fasthttp.Request, res *fasthttp.Response, deadline time.Time) error
}
