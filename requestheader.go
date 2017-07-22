package nrfasthttp

import "github.com/valyala/fasthttp"

type RequestHeader struct {
	*fasthttp.RequestHeader
}

func (r RequestHeader) Get(name string) string {
	var result string

	// This if by far the worst implementation possible
	// Unfortunately, it's the only way we can do it given
	// the current fasthttp interface
	r.RequestHeader.VisitAll(func(key, value []byte) {
		if string(key) == name {
			result = string(value)
		}
	})

	return result
}
