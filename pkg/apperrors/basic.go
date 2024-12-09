package apperrors

import "fmt"

type ErrorContent struct {
	Code    string
	Message string
}

type ValidationError struct {
	Code    string
	Message string
	Field   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewValidationError(content ErrorContent, field string) *ValidationError {
	return &ValidationError{
		Code:    content.Code,
		Message: content.Message,
		Field:   field,
	}
}

type NotFoundError struct {
	Code    string
	Message string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewNotFound(content ErrorContent) *NotFoundError {
	return &NotFoundError{
		Code:    content.Code,
		Message: content.Message,
	}
}

type PermissionsError struct {
	Code    string
	Message string
}

func (e *PermissionsError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewPermissionsError(content ErrorContent) *PermissionsError {
	return &PermissionsError{
		Code:    content.Code,
		Message: content.Message,
	}
}

type UnauthorizedError struct {
	Code    string
	Message string
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewUnauthorized(content ErrorContent) *UnauthorizedError {
	return &UnauthorizedError{
		Code:    content.Code,
		Message: content.Message,
	}
}

type InternalError struct {
	Code    string
	Message string
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewInternalError(content ErrorContent) *InternalError {
	return &InternalError{
		Code:    content.Code,
		Message: content.Message,
	}
}
