package middleware

import (
	"errors"
	"net/http"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/carrizoaagustin/cv-online/pkg/apperrors"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func camelCaseToSnakeCase(s string) string {
	var result []rune
	for i, char := range s {
		if unicode.IsUpper(char) && i > 0 {
			result = append(result, '_', unicode.ToLower(char))
		} else {
			result = append(result, char)
		}
	}
	return strings.ToLower(string(result))
}

func extractValidationErrors(err error) gin.H {
	validationErrors := make(gin.H)

	u, ok := err.(interface {
		Unwrap() []error
	})

	var validationErr *apperrors.ValidationError

	if !ok && errors.As(err, &validationErr) {
		validationErrors[camelCaseToSnakeCase(validationErr.Field)] = ErrorResponse{
			Code:    validationErr.Code,
			Message: validationErr.Message,
		}
	}

	for _, errorCustom := range u.Unwrap() {
		if errors.As(errorCustom, &validationErr) {
			validationErrors[camelCaseToSnakeCase(validationErr.Field)] = ErrorResponse{
				Code:    validationErr.Code,
				Message: validationErr.Message,
			}
		}
	}
	return validationErrors
}

func extractPlaygroundValidationErrors(err error) gin.H {
	playgroundErrors := make(gin.H)

	var playgroundValidation validator.ValidationErrors

	if errors.As(err, &playgroundValidation) {
		for _, fieldError := range playgroundValidation {
			response := getPlaygroundMapper(fieldError.Tag(), fieldError.Param())

			playgroundErrors[camelCaseToSnakeCase(fieldError.Field())] = ErrorResponse{
				Code:    response.Code,
				Message: response.Message,
			}
		}
	}

	return playgroundErrors
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			var validationErrors gin.H
			var playgroundErrors gin.H
			for _, err := range c.Errors {
				var validationErr *apperrors.ValidationError
				var notFoundErr *apperrors.NotFoundError
				var permissionsErr *apperrors.PermissionsError
				var unauthorizedErr *apperrors.UnauthorizedError
				var internalErr *apperrors.InternalError
				var playgroundValidation validator.ValidationErrors
				switch {
				case errors.As(err.Err, &notFoundErr):
					c.AbortWithStatusJSON(http.StatusNotFound, ErrorResponse{
						Code:    notFoundErr.Code,
						Message: notFoundErr.Message,
					})
					return

				case errors.As(err.Err, &permissionsErr):
					c.AbortWithStatusJSON(http.StatusForbidden, ErrorResponse{
						Code:    permissionsErr.Code,
						Message: permissionsErr.Message,
					})
					return

				case errors.As(err.Err, &unauthorizedErr):
					c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
						Code:    unauthorizedErr.Code,
						Message: unauthorizedErr.Message,
					})
					return

				case errors.As(err.Err, &internalErr):
					// ADD LOG
					c.AbortWithStatusJSON(http.StatusForbidden, ErrorResponse{
						Code:    internalErr.Code,
						Message: internalErr.Message,
					})
					return

				case errors.As(err.Err, &playgroundValidation):
					playgroundErrors = extractPlaygroundValidationErrors(err.Err)

				case errors.As(err.Err, &validationErr):
					validationErrors = extractValidationErrors(err.Err)

				default:
					// ADD LOG
					c.AbortWithStatusJSON(http.StatusForbidden, ErrorResponse{
						Code:    "UNKNOWN_ERROR",
						Message: "unexpected error",
					})
					return
				}
			}

			if len(playgroundErrors) > 0 {
				c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
					Code:    "VALIDATION_ERROR",
					Message: "Invalid fields. Please check details.",
					Details: playgroundErrors,
				})
				return
			}

			if len(validationErrors) > 0 {
				c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
					Code:    "VALIDATION_ERROR",
					Message: "Invalid fields. Please check details.",
					Details: validationErrors,
				})
				return
			}
		}
	}
}
