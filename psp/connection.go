package psp

import "net/http"

type Connection interface {
	Do(Request) (body []byte, headers http.Header, err error)
}
