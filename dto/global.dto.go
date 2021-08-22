package dto

type Response struct {
	StatusCode int         `json:"statusCode"`
	StatusDesc string      `json:"statusDesc"`
	Content    interface{} `json:"content"`
}
