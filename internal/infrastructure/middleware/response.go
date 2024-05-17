package middleware

import (
	"github.com/go-playground/validator/v10"
	error2 "service/internal/infrastructure/error"
)

// Response response.
type Response[T any] struct {
	Success      bool   `json:"success"`
	ErrorCode    int64  `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Data         T      `json:"data"`
}

func success[T any](data T) Response[T] {
	return Response[T]{
		Success:      true,
		ErrorCode:    0,
		ErrorMessage: "",
		Data:         data,
	}
}

func ofServiceError(err error2.ServiceError) Response[struct{}] {
	return Response[struct{}]{
		Success:      false,
		ErrorCode:    err.ErrorCode,
		ErrorMessage: err.ErrorMessage,
		Data:         struct{}{},
	}
}

func ofValidationError(err validator.ValidationErrors) Response[struct{}] {
	return Response[struct{}]{
		Success:      false,
		ErrorCode:    400,
		ErrorMessage: err.Error(),
		Data:         struct{}{},
	}
}
