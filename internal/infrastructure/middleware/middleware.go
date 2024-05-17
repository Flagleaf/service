package middleware

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/gin-gonic/gin"
	error2 "service/internal/infrastructure/error"
)

func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var serviceError *error2.ServiceError
			if errors.As(err, &serviceError) {
				c.JSON(http.StatusInternalServerError, ofServiceError(*serviceError))
				return
			}

			var validationErrors *validator.ValidationErrors
			if errors.As(err, &validationErrors) {
				c.JSON(http.StatusBadRequest, ofValidationError(*validationErrors))
				return
			}

			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if c.Writer.Status() == 0 {
			c.Writer.WriteHeader(http.StatusOK)
		}

		if c.Writer.Status() >= http.StatusOK && c.Writer.Status() < http.StatusMultipleChoices {
			data, exists := c.Get("response")
			if exists {
				c.JSON(http.StatusOK, success(data))
				return
			}
		}
	}
}
