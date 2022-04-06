package models

// JSONResponses : json response
type JSONResponses struct {
	Code   int         `json:"code"`
	Status interface{} `json:"status"`
	Data   interface{} `json:"data"`
}

// JSONResponsesPagination : json with pagination
type JSONResponsesPagination struct {
	Code   int         `json:"code"`
	Status interface{} `json:"status"`
	Data   interface{} `json:"data"`
	Meta   interface{} `json:"meta"`
}
