package http

import "net/http"

type HttpMethod int

const (
	GET HttpMethod = iota
	POST
	PUT
	DELETE
	PATCH
)

var httpMethods = map[HttpMethod]string{
	GET:    http.MethodGet,
	POST:   http.MethodPost,
	PUT:    http.MethodPut,
	DELETE: http.MethodDelete,
	PATCH:  http.MethodPatch,
}

func (method HttpMethod) Value() string {

	if value, ok := httpMethods[method]; ok {
		return value
	}

	return ""

}

type HttpResult struct {
	Code int
	Data string
}
