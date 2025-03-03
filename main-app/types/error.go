package types

type NotFoundError struct {
	Message string `json:"message"`
}
type UnauthorizedError struct {
	Message string `json:"message"`
}
type BadRequestError struct {
	Message string `json:"message"`
}

type InternalServerError struct {
	Message string `json:"message"`
}

func (e *NotFoundError) Error() string {
	return e.Message
}
func (e *UnauthorizedError) Error() string {
	return e.Message
}
func (e *BadRequestError) Error() string {
	return e.Message
}
func (e *InternalServerError) Error() string {
	return e.Message
}
