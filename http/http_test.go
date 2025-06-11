package http

import (
	"fmt"
	"testing"
)

func TestGetRequest(t *testing.T) {

	headers := make(map[string]string)
	headers["Authorization"] = "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IkY1Qzc1QUEwQTE1M0M1OTI4NzE0QTk5OEJFRkJFQkQ5IiwidHlwIjoiYXQrand0In0.eyJuYmYiOjE3NDk2MjIyMjAsImV4cCI6MTc2ODU0MzgyMCwiaXNzIjoiaHR0cDovLzEwLjQwLjEyOS45Mzo1Mjk1IiwiY2xpZW50X2lkIjoidGVzdCIsImp0aSI6IkI0NzgwMEMyODMzMjJCRDYxQjhBQTk1RTZFQTJDQzBFIiwiaWF0IjoxNzQ5NjIyMjIwLCJzY29wZSI6WyJhZHRlc3QiLCJhbG9uZGEiLCJiYW5jYXNzdXJhbmNlIiwiYndiIiwiYndiX3Bvc3RwYWlkIiwiY3VuaW1hIiwiZHJ0c3MiLCJlYXp5SW50ZWdyYXRpb25zIiwiZWRnZXRlY2hfcGF5IiwiZWt5YyIsImZsZXhjdWJlIiwiZ290diIsImh1YmdhdGV3YXkiLCJrYWNoYW5ndSIsImtodXNhIiwibHdiIiwibHdiX25ld19wb3N0cGFpZCIsIm1haWMiLCJtYW1hX21vbmV5IiwibXVrdXJ1IiwibXVsdGljaG9pY2UiLCJuZWVmIiwibm90aWZpY2F0aW9uZ2F0ZXdheSIsIm9yYWNsZWZsZXhjdWJlIiwicGFmdXBpIiwidGMiXX0.WDljJzuzg20piJaRLKdyrytM7y_1I3eR3FYPHbokdLB1VQtMIjWPJHTmO2C6C8muTS8VYI5XRquPj7SO3V0GnEjbPm5oxUHe69Ybk9GJL-UyHKR4gccl5CVsY17z5CeZWhvQvCMQNQZ83MDMBPKsKIkSwHugC7J1rk2tzVKbm-Gse0Yplm63Mg23EI-Ks8GlrNwl--ETDtszKGzCu_eo5UUbZ3k-8hZ9FcL-pFQDgwa5LPUJ2RcKbNDjHUUs5u5_4wDUr-tFl503WXjBqkttITgt_1GT5GVJMHQK7dVmMvXkIL2TtIXTQqxitHwkz55y2vXQ-6nPLU5llTe3L6746Q"

	response, err := GET("http://nbsdevtest:5145/api/main/CheckCBSStatus", headers)

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
	headers["Authorization"] = "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IkY1Qzc1QUEwQTE1M0M1OTI4NzE0QTk5OEJFRkJFQkQ5IiwidHlwIjoiYXQrand0In0.eyJuYmYiOjE3NDk2MjI0NzYsImV4cCI6MTc2ODU0NDA3NiwiaXNzIjoiaHR0cDovLzEwLjQwLjEzMC4yNTo1ODc1IiwiY2xpZW50X2lkIjoibmJzZGV2cyIsImp0aSI6IkQ1NjY2QTQwREJENTVBQjJEMDIxODg1NzU2Q0U3REM0IiwiaWF0IjoxNzQ5NjIyNDc2LCJzY29wZSI6WyJiYW5jYXNzdXJhbmNlIiwiYndiIiwiZmxleGN1YmUiLCJnb3R2IiwiaHViZ2F0ZXdheSIsImthY2hhbmd1IiwibXVrdXJ1IiwibXVsdGljaG9pY2UiLCJuZWVmIiwibm90aWZpY2F0aW9uZ2F0ZXdheSIsInBhZnVwaSJdfQ.b8RjK2P0HZinS_on6X0Ei6C-ERp3Jc5Ib8Db-hkaUFFyA_NNf6ElBBFRfCQHOTKxME0XB_ZRzfU8DZ5x7M6ROrcpW2R5NiECQa-7LT-VE7m0ijM8VwOo8MyAlfNMYJRDkramUm-g0w0mYIfhI2oQddKkt-MhGk8jMDY1OyjYNoeBde9zzhtgxcrMT0aaG4ArbtsQnpzAo7W2a_pPIob9GsJCt8VmZnpgWbzzMSBZ_PmSQ-_0wA9arZ0xsyKnIKZ69X8lzwywdwee1uoXExpnaSzbpZ-gGRljuO-mBNlW2h73O5B6GUjw814mVsHKoBY_w1lOGfFnPp1MMw_N4nyKXw"

	type ValidateAccountRequest struct {
		AccountNumber string `json:"accountNumber"`
	}

	request := ValidateAccountRequest{AccountNumber: "22904248"}

	body, err := JSONBody(request)

	if err != nil {
		t.Errorf("HTTP Request not sent, error encountered: %s", err.Error())
	}

	response, err := POST("http://10.40.129.57:84/flex/account/information", body, headers)

	if err != nil {
		t.Errorf("HTTP Request not sent, error encountered: %s", err.Error())
	}

	if response != nil {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Response)
	}

}
