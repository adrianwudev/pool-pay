package util

import (
	"log"
	"pool-pay/internal/constants"
	"runtime/debug"
)

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ApiError) Error() string {
	return e.Message
}

func SetApiError(code int) *ApiError {
	return &ApiError{
		Code:    code,
		Message: constants.ErrorCodeMessageMap[code],
	}
}

func SetDefaultApiError(err error) *ApiError {
	log.Println(err)
	log.Println(string(debug.Stack()))
	return &ApiError{
		Code:    constants.ERRORCODE_OTHERS,
		Message: err.Error(),
	}
}

func GetApiError(err error) *ApiError {
	log.Println(err)
	log.Println(string(debug.Stack()))
	if err != nil {
		if apiErr, ok := err.(*ApiError); ok {
			return apiErr
		} else {
			log.Printf("error can't convert to ApiError: %s\n", err)
			return nil
		}
	}
	log.Printf("error is nil: %s\n", err)
	return nil
}
