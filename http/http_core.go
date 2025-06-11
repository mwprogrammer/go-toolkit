package http

import (
	"io"
	"net/http"
	"time"
)

type HttpResponse struct {
	StatusCode int
	Response   string
}

type HttpRequestBody struct {
	Reader      io.Reader
	ContentType string
}

func Send(url, method string, body *HttpRequestBody, headers map[string]string) (*HttpResponse, error) {

	if headers == nil {
		headers = make(map[string]string)
	}

	if body != nil && body.ContentType != "" {
		if _, exists := headers["Content-Type"]; !exists {
			headers["Content-Type"] = body.ContentType
		}
	}

	var reader io.Reader

	if body != nil {
		reader = body.Reader
	}

	return SendRequest(url, method, reader, headers)

}

func SendRequest(url, method string, body io.Reader, headers map[string]string) (*HttpResponse, error) {

	request, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{Timeout: 30 * time.Second}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	if err != nil {
		content = []byte{}
	}

	result := new(HttpResponse)
	result.StatusCode = response.StatusCode
	result.Response = string(content)

	return result, nil

}
