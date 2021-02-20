package utils

// APIResponse struct for all api responses
type APIResponse struct {
	Error *string     `json:"error"`
	Data  interface{} `json:"data"`
}
