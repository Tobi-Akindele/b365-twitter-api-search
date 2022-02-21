package controllers

import (
	"bet365/dtos"
	"encoding/json"
	"net/http"
)

func Ping(rw http.ResponseWriter, _ *http.Request) {
	response := dtos.Response{
		Message: "Bet365-Twitter-bot up!",
	}
	_ = json.NewEncoder(rw).Encode(response)
}
