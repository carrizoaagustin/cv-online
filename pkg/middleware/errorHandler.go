package middleware

import (
	"net/http"
	"unicode"

	"github.com/gin-gonic/gin"

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
	return string(result)
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			validationErrors := make(gin.H, 0)

			for _, err := range c.Errors {
				switch e := err.Err.(type) {
				case *apperrors.ValidationError:
					validationErrors[camelCaseToSnakeCase(e.Field)] = ErrorResponse{
						Code:    e.Code,
						Message: e.Message,
					}

				case *apperrors.NotFoundError:
					c.AbortWithStatusJSON(http.StatusNotFound, ErrorResponse{
						Code:    e.Code,
						Message: e.Message,
					})

				case *apperrors.PermissionsError:
					c.AbortWithStatusJSON(http.StatusForbidden, ErrorResponse{
						Code:    e.Code,
						Message: e.Message,
					})

				case *apperrors.UnauthorizedError:
					c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
						Code:    e.Code,
						Message: e.Message,
					})

				case *apperrors.InternalError:
					// ADD LOG
					c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{
						Code:    e.Code,
						Message: e.Message,
					})

				default:
					// ADD LOG
					c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{
						Code:    "UNKNOWN_ERROR",
						Message: "u.",
					})
				}
			}

			if len(validationErrors) > 0 {
				c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
					Code:    "VALIDATION_ERROR",
					Message: "Invalid fields. Please check details.",
					Details: validationErrors,
				})
			}
		}
	}
}
