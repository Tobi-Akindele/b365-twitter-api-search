package controllers

import (
	"bet365/services"
	"bet365/utils"
	"encoding/json"
	"net/http"
)

func TwitterBot(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	queryString := r.URL.Query().Get("q")
	if utils.IsEmptyString(queryString) {
		response := utils.ComposeResponse(nil, "Query string is required")
		_ = json.NewEncoder(rw).Encode(response)
		return
	}
	queryResponse, err := services.QueryTwitterAPI(queryString)
	if err != nil {
		response := utils.ComposeResponse(nil, err.Error())
		_ = json.NewEncoder(rw).Encode(response)
		return
	}
	_ = json.NewEncoder(rw).Encode(queryResponse)
}
