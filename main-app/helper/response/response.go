package response

import "github.com/revandpratama/reflect/types"

type ResponseWithData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithoutData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewResponse(response *types.ResponseParams) any {
	var res any
	var status string

	if response.StatusCode > 400 {
		status = "error"
	} else {
		status = "success"
	}

	if response.Data != nil {
		res = &ResponseWithData{
			Status:  status,
			Message: response.Message,
			Data:    response.Data,
		}
	} else {
		res = &ResponseWithoutData{
			Status:  status,
			Message: response.Message,
		}
	}

	return res
}
