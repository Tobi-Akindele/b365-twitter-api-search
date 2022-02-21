package utils

import (
	"bet365/dtos"
	"log"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

func IsEmptyString(str string) bool {
	str = strings.Trim(str, "")
	return str == "" || str == "nil" || str == "null"
}

func ComposeResponse(data interface{}, message string) dtos.Response {
	return dtos.Response{
		Data:    data,
		Message: message}
}
