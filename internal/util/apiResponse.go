package util

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(message string, data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(err error) APIResponse {
	return APIResponse{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}
}
