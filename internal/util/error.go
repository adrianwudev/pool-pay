package util

import (
	"log"
	"pool-pay/internal/constants"
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
	return &ApiError{
		Code:    constants.ERRORCODE_OTHERS,
		Message: err.Error(),
	}
}

func GetApiError(err error) *ApiError {
	if err != nil {
		if apiErr, ok := err.(*ApiError); ok {
			return apiErr
		} else {
			log.Fatalf("error can't convert to ApiError: %s\n", err)
			return nil
		}
	}
	log.Printf("error is nil: %s\n", err)
	return nil
}
