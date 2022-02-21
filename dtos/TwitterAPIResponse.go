package dtos

type TweetData struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type TwitterAPIResponse struct {
	Data []TweetData `json:"data"`
}
