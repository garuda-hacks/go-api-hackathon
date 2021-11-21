package response

type ErrorResponse struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
}
type EmptyObj struct{}
