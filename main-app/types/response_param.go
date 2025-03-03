package types

type ResponseParams struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}