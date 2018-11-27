package nrfasthttp

import fasthttp "gopkg.in/valyala/fasthttp.v20160316"

type ResponseHeader struct {
	*fasthttp.ResponseHeader
}

func (r ResponseHeader) Add(name, value string) {
	r.ResponseHeader.Add(name, value)
}

func (r ResponseHeader) Get(name string) string {
	var result string

	// This if by far the worst implementation possible
	// Unfortunately, it's the only way we can do it given
	// the current fasthttp interface
	r.ResponseHeader.VisitAll(func(key, value []byte) {
		if string(key) == name {
			result = string(value)
		}
	})

	return result
}
