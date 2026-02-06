package types

import "fmt"

type ErrorRequest struct {
	Message string
	Code    int
}

func NewErrorRequest(message string, code ...int) *ErrorRequest {
	statusCode := 400
	if len(code) > 0 {
		statusCode = code[0]
	}
	return &ErrorRequest{
		Message: message,
		Code:    statusCode,
	}
}

func (e *ErrorRequest) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}
