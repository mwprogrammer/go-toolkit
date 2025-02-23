package http

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

func Get[T interface{}](url string, token string, logger *slog.Logger) (T, bool) {

	is_success := true

	logger.Info("HTTP REQUEST", "URL", url, "BODY", "")

	var data T
	response, err := http.Get(url)

	if err != nil {

		logger.Error(err.Error())
		is_success = false

		return data, is_success
	}

	defer response.Body.Close()

	output, err := io.ReadAll(response.Body)

	defer response.Body.Close()

	logger.Info("HTTP RESPONSE", "URL", url, "STATUS CODE", response.StatusCode, "RESPONSE", output)

	if (err == nil) {
		
		cbs_err := json.Unmarshal([]byte(output), &data)

		if cbs_err != nil {
			logger.Error(cbs_err.Error())
			is_success = false
		}
		
	}
	
	return data, is_success


}

func Post[T interface{}](url string, body []byte, token string, logger *slog.Logger) (T, bool){

	is_success := true

	logger.Info("HTTP REQUEST", "URL", url, "BODY", string(body))

	var data T
	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {

		logger.Error(err.Error())
		is_success = false

		return data, is_success
	}

	defer response.Body.Close()

	output, err := io.ReadAll(response.Body)

	logger.Info("HTTP RESPONSE", "URL", url, "STATUS CODE", response.StatusCode, "RESPONSE", output)

	if(err == nil) {
		
		cbs_err := json.Unmarshal([]byte(output), &data)

		if cbs_err != nil {
			logger.Error(cbs_err.Error())
			is_success = false
		}
		
	}
	
	return data, is_success
}