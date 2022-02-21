package services

import (
	"bet365/config"
	"bet365/dtos"
	"bet365/utils"
	"encoding/json"
	"strings"
)

func QueryTwitterAPI(queryString string) (dtos.QueryResponse, error) {

	apiUrl := config.GetAppProperty("TWITTER_SEARCH_API")
	accessToken := config.GetAppProperty("ACCESS_TOKEN")
	maxResults := config.GetAppProperty("MAX_RESULTS")
	requestParams := map[string]string{
		utils.QUERY:      queryString,
		utils.MAX_RESULT: maxResults,
	}
	requestHeaders := map[string]string{
		utils.AUTHORIZATION: "Bearer " + accessToken,
	}

	stringResponse, err := GetRequest(apiUrl, requestParams, requestHeaders)
	if err != nil {
		return dtos.QueryResponse{}, err
	}

	response := dtos.TwitterAPIResponse{}
	_ = json.Unmarshal([]byte(stringResponse), &response)
	var resultCount int
	var avgWordCount int
	if response.Data != nil {
		resultCount = len(response.Data)
		wordCountTotal := 0
		for i := 0; i < resultCount; i++ {
			wordCountTotal += len(strings.Fields(response.Data[i].Text))
		}
		avgWordCount = wordCountTotal / resultCount
	}

	return dtos.QueryResponse{
		SearchTerm:   queryString,
		ResultCount:  resultCount,
		AvgWordCount: avgWordCount,
	}, nil
}
