package http

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

func SendRequest(method HttpMethod, URL string, json []byte) (*HttpResult, error) {

	body := bytes.NewReader(json)

	request, err := http.NewRequest(method.Value(), URL, body)

	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	if err != nil {
		content = []byte{}
	}

	result := new(HttpResult)
	result.Code = response.StatusCode
	result.Data = string(content)

	return result, nil

}

func SendRequestWithNoBody(method HttpMethod, URL string) (*HttpResult, error) {

	request, err := http.NewRequest(method.Value(), URL, nil)

	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	response.Body.Close()

	content, err := io.ReadAll(response.Body)

	if err != nil {
		content = []byte{}
	}

	result := new(HttpResult)
	result.Code = response.StatusCode
	result.Data = string(content)

	return result, nil

}

func SendRequestWithFormData(method HttpMethod, URL string, form map[string]string) (*HttpResult, error) {

	fields := url.Values{}

	for key, value := range form {
		fields.Set(key, value)
	}

	body := bytes.NewBufferString(fields.Encode())

	request, err := http.NewRequest(method.Value(), URL, body)

	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	if err != nil {
		content = []byte{}
	}

	result := new(HttpResult)
	result.Code = response.StatusCode
	result.Data = string(content)

	return result, nil

}
