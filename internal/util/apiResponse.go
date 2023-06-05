package util

import "time"

type APIResponse struct {
	Success   bool        `json:"success"`
	Timestamp string      `json:"timestamp" example:"2021-07-29T07:23:47Z"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	ErrorCode int         `json:"errorCode"`
}

func NewSuccessResponse(message string, data interface{}) APIResponse {
	return APIResponse{
		Success:   true,
		Timestamp: time.Now().Format(time.RFC3339),
		Message:   message,
		Data:      data,
		ErrorCode: 0,
	}
}

func NewErrorResponse(err error, errCode int) APIResponse {
	return APIResponse{
		Success:   false,
		Timestamp: time.Now().Format(time.RFC3339),
		Message:   err.Error(),
		Data:      nil,
		ErrorCode: errCode,
	}
}
