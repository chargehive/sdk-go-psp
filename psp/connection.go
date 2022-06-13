package psp

import "net/http"

type Connection interface {
	Do(Request) (*http.Response, error)
}
