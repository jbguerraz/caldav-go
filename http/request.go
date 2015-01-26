package http

import (
	"github.com/taviti/caldav-go/utils"
	"io"
	"net/http"
)

// an HTTP request object
type Request http.Request

// downcasts the request to the native HTTP interface
func (r *Request) Native() *http.Request {
	return (*http.Request)(r)
}

// creates a new HTTP request object
func NewRequest(method string, urlstr string, body ...io.ReadCloser) (*Request, error) {
	var err error
	var r = new(http.Request)
	if len(body) > 0 && body[0] != nil {
		r, err = http.NewRequest(method, urlstr, body[0])
	} else {
		r, err = http.NewRequest(method, urlstr, nil)
	}
	if err != nil {
		return nil, utils.NewError(NewRequest, "unable to create request", urlstr, err)
	} else {
		return (*Request)(r), nil
	}
}