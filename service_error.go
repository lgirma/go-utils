package utils

import (
	"fmt"
	"strconv"
)

type BusinessError struct {
	Code   string
	Detail any
}

func (err *BusinessError) Error() string {
	result := err.Code
	if err.Detail != nil {
		result = fmt.Sprintf("%s: %v", result, err.Detail)
	}
	return result
}

func NewBusinessError(code string, details ...any) *BusinessError {
	var detail any
	if len(details) > 0 {
		detail = details[0]
	}
	return &BusinessError{
		Code:   code,
		Detail: detail,
	}
}

func NewBusinessErrorFromErrNo(errNo int, details ...any) *BusinessError {
	var detail any
	if len(details) > 0 {
		detail = details[0]
	}
	return &BusinessError{
		Code:   strconv.Itoa(errNo),
		Detail: detail,
	}
}

func IsBusinessError(err error) bool {
	_, ok := err.(*BusinessError)
	return ok
}