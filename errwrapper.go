package nrfasthttp

import "fmt"

type errWrapper struct {
	err interface{}
}

func (e errWrapper) Error() string {
	return fmt.Sprintf("generic error %v", e.err)
}
