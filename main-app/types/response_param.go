package types

type ResponseParams struct {
	StatusCode int         `json:"status"`
	Message    string      `json:"message"`
	Data       any         `json:"data"`
	Errors     any         `json:"errors"`
	Pagination *Pagination `json:"pagination"`
}
