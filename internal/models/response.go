package models

type Response struct {
	Success bool `json:"success"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}