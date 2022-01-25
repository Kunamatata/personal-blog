package types

//APIResponse is the response format for API responses
type APIResponse struct {
	Status  string      `json:"status"`
	Message *string     `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
