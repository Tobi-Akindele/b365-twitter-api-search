package services

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func resolveRequestParams(req *http.Request, params map[string]string) {
	requestParams := req.URL.Query()
	for key, value := range params {
		requestParams.Add(key, value)
	}
	req.URL.RawQuery = requestParams.Encode()
}

func resolveRequestHeaders(req *http.Request, requestHeaders map[string]string) {
	for key, value := range requestHeaders {
		req.Header.Set(key, value)
	}
}

func GetRequest(apiUrl string, params map[string]string, requestHeaders map[string]string) (string, error) {

	req, _ := http.NewRequest(http.MethodGet, apiUrl, nil)
	resolveRequestParams(req, params)
	resolveRequestHeaders(req, requestHeaders)

	client := &http.Client{}
	res, err := client.Do(req)
	logError(err)

	defer handleResponseBody(res.Body)

	var responseBody string
	var errorMessage error
	if res.StatusCode == http.StatusOK {
		bodyInBytes, _ := io.ReadAll(res.Body)
		responseBody = string(bodyInBytes)
	} else {
		errorMessage = errors.New(fmt.Sprintf("Status code returned is %v ", res.Status))
	}

	return responseBody, errorMessage
}

func logError(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

func handleResponseBody(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Println(err)
	}
}
