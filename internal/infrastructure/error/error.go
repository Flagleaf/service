package error

import (
	"fmt"
	"strconv"
)

type ServiceError struct {
	ErrorCode    int64
	ErrorMessage string
}

func NewServiceError(code int64, message string) error {
	return &ServiceError{
		ErrorCode:    code,
		ErrorMessage: message,
	}
}

func (s *ServiceError) Error() string {
	return fmt.Sprintf("[errorCode=%s, errorMessage=%s]",
		strconv.FormatInt(s.ErrorCode, 0), s.ErrorMessage)
}
