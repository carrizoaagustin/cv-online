package middleware

import "fmt"

//nolint:gochecknoglobals //i only can access this var in middleware package
var playgroundMapper = map[string]struct {
	Code    string
	Message string
}{
	"required": {
		Code:    "REQUIRED",
		Message: "This field is required.",
	},
	"min": {
		Code:    "MIN_VALUE",
		Message: "The value must be at least %s.",
	},
	"email": {
		Code:    "INVALID_EMAIL",
		Message: "The value must be a valid email.",
	},
}

func getPlaygroundMapper(tag string, param string) struct {
	Code    string
	Message string
} {
	if value, ok := playgroundMapper[tag]; ok {
		var message string
		if param == "" {
			message = value.Message
		} else {
			message = fmt.Sprintf(value.Message, param)
		}

		return struct {
			Code    string
			Message string
		}{
			Code:    value.Code,
			Message: message,
		}
	}
	return struct {
		Code    string
		Message string
	}{
		Code:    "UNKNOWN_VALIDATION_ERROR",
		Message: "Unknown validation error.",
	}
}
