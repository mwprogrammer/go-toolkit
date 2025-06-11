package http

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/url"
	"strings"
)

func NoBody() *HttpRequestBody {
	return &HttpRequestBody{Reader: nil}
}

func JSONBody(data any) (*HttpRequestBody, error) {

	if data == nil {
		return NoBody(), nil
	}

	jsonBytes, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return &HttpRequestBody{
		Reader:      bytes.NewBuffer(jsonBytes),
		ContentType: "application/json",
	}, nil

}

func XMLBody(data any) (*HttpRequestBody, error) {

	if data == nil {
		return NoBody(), nil
	}

	xmlBytes, err := xml.Marshal(data)

	if err != nil {
		return nil, err
	}

	return &HttpRequestBody{
		Reader:      bytes.NewBuffer(xmlBytes),
		ContentType: "application/xml",
	}, nil

}

func FormBody(data map[string]string) *HttpRequestBody {

	form := url.Values{}

	for key, value := range data {
		form.Set(key, value)
	}

	encoded := form.Encode()

	return &HttpRequestBody{
		Reader:      strings.NewReader(encoded),
		ContentType: "application/x-www-form-urlencoded",
	}
}
