package http

import "net/http"

func GET(url string, headers map[string]string) (*HttpResponse, error) {
	return Send(url, http.MethodGet, NoBody(), headers)
}

func POST(url string, body *HttpRequestBody, headers map[string]string) (*HttpResponse, error) {
	return Send(url, http.MethodPost, body, headers)
}

func PUT(url string, body *HttpRequestBody, headers map[string]string) (*HttpResponse, error) {
	return Send(url, http.MethodPut, body, headers)
}

func DELETE(url string, body *HttpRequestBody, headers map[string]string) (*HttpResponse, error) {
	return Send(url, http.MethodDelete, body, headers)
}

func PATCH(url string, body *HttpRequestBody, headers map[string]string) (*HttpResponse, error) {
	return Send(url, http.MethodPatch, body, headers)
}
