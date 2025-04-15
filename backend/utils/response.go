package utils

type Response struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message interface{}) Response {
	return Response{
		Success: false,
		Message: message,
	}
}
