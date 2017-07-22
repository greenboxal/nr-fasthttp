package nrfasthttp

import (
	"net/url"

	"github.com/newrelic/go-agent/http"
	"github.com/valyala/fasthttp"
)

type Request struct {
	*fasthttp.Request
}

func (r Request) Header() http.Header {
	return RequestHeader{&r.Request.Header}
}

func (r Request) Method() string {
	return string(r.Request.Header.Method())
}

func (r Request) URL() *url.URL {
	uri := r.Request.URI()

	return &url.URL{
		Scheme:   string(uri.Scheme()),
		Path:     string(uri.Path()),
		Host:     string(uri.Host()),
		RawQuery: string(uri.QueryString()),
	}
}
