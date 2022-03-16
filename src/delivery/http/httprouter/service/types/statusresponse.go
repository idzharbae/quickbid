package types

type StatusResponse struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}
