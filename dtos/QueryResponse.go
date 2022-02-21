package dtos

type QueryResponse struct {
	SearchTerm   string `json:"searchTerm"`
	ResultCount  int    `json:"resultCount"`
	AvgWordCount int    `json:"avgWordCount"`
}
