package handlers

type Response struct {
	Success bool `json:"success"`
	Code string `json:"code"`
	Message string `json:"message"`
	Data any `json:"data"`
}