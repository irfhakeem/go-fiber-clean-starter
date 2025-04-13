package utils

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
	Meta    any    `json:"meta,omitempty"`
}

func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func FailedResponse(message string, err string) Response {
	return Response{
		Status:  false,
		Message: message,
		Error:   err,
	}
}
