package http

import (
	"fmt"
	"testing"
)

func TestGetRequest(t *testing.T) {

	headers := make(map[string]string)
	headers["Authorization"] = "Bearer token..."

	response, err := GET("url", headers)

	if err != nil {
		t.Errorf("HTTP Request not sent, error encountered: %s", err.Error())
	}

	if response != nil {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Response)
	}
}

func TestPostRequest(t *testing.T) {

	headers := make(map[string]string)
	headers["Authorization"] = "Bearer token..."

	type SampleRequest struct {
		SampleProperty string `json:"accountNumber"`
	}

	request := SampleRequest{SampleProperty: "value"}

	body, err := JSONBody(request)

	if err != nil {
		t.Errorf("HTTP Request not sent, error encountered: %s", err.Error())
	}

	response, err := POST("url", body, headers)

	if err != nil {
		t.Errorf("HTTP Request not sent, error encountered: %s", err.Error())
	}

	if response != nil {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Response)
	}

}
