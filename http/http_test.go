package http

import (
	"fmt"
	"testing"
)

func TestGetRequest(t *testing.T) {

	result, err := SendRequestWithNoBody(GET, "http://nbsdevtest:5145/api/main/CheckCBSStatus")

	if err != nil {
		t.Errorf("HTTP Request not sent, error encountered: %s", err.Error())
	}

	if result != nil {
		fmt.Println(result.Data)
	}
}

func TestPostRequest(t *testing.T) {

	request := `{"accountNumber": "22904248"}`

	result, err := SendRequest(POST, "http://10.40.129.57:84/flex/account/information", []byte(request))

	if err != nil {
		t.Errorf("HTTP Request not sent, error encountered: %s", err.Error())
	}

	if result != nil {
		fmt.Println(result.Code)
	}

}
