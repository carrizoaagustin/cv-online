package middleware

import "fmt"

//nolint:gochecknoglobals //i only can access this var in middleware package
var playgroundMapper = map[string]struct {
	Code    string
	Message string
}{
	"required": {
		Code:    "REQUIRED_ERROR",
		Message: "This field is required.",
	},
	"min": {
		Code:    "MIN_VALUE_ERROR",
		Message: "The value must be at least %s.",
	},
	"max": {
		Code:    "MAX_VALUE_ERROR",
		Message: "The value must be at most %s.",
	},
	"len": {
		Code:    "INVALID_LENGTH",
		Message: "The value must have a length of %s.",
	},
	"email": {
		Code:    "INVALID_EMAIL",
		Message: "The value must be a valid email.",
	},
	"url": {
		Code:    "INVALID_URL",
		Message: "The value must be a valid URL.",
	},
	"uuid": {
		Code:    "INVALID_UUID",
		Message: "The value must be a valid UUID.",
	},
	"alpha": {
		Code:    "INVALID_ALPHA",
		Message: "The value must contain only alphabetic characters.",
	},
	"alphanum": {
		Code:    "INVALID_ALPHANUM",
		Message: "The value must contain only alphanumeric characters.",
	},
	"numeric": {
		Code:    "INVALID_NUMERIC",
		Message: "The value must be numeric.",
	},
	"positive": {
		Code:    "INVALID_POSITIVE",
		Message: "The value must be positive.",
	},
	"negative": {
		Code:    "INVALID_NEGATIVE",
		Message: "The value must be negative.",
	},
	"iso8601": {
		Code:    "INVALID_ISO8601",
		Message: "The value must be a valid ISO 8601 date/time.",
	},
	"timezone": {
		Code:    "INVALID_TIMEZONE",
		Message: "The value must be a valid timezone.",
	},
	"datetime": {
		Code:    "INVALID_DATETIME",
		Message: "The value must be a valid datetime.",
	},
	"required_if": {
		Code:    "REQUIRED_IF",
		Message: "This field is required when %s is present.",
	},
	"required_unless": {
		Code:    "REQUIRED_UNLESS",
		Message: "This field is required unless %s is present.",
	},
	"excluded_with": {
		Code:    "EXCLUDED_WITH",
		Message: "This field is excluded when %s is present.",
	},
	"excluded_with_all": {
		Code:    "EXCLUDED_WITH_ALL",
		Message: "This field is excluded when any of %s are present.",
	},
	"not_empty": {
		Code:    "NOT_EMPTY",
		Message: "The value must not be empty.",
	},
	"contains": {
		Code:    "CONTAINS",
		Message: "The value must contain %s.",
	},
	"starts_with": {
		Code:    "STARTS_WITH",
		Message: "The value must start with %s.",
	},
	"ends_with": {
		Code:    "ENDS_WITH",
		Message: "The value must end with %s.",
	},
	"excluded": {
		Code:    "EXCLUDED",
		Message: "The value must not be in the list of excluded values.",
	},
	"lte": {
		Code:    "LTE",
		Message: "The value must be less than or equal to %s.",
	},
	"gte": {
		Code:    "GTE",
		Message: "The value must be greater than or equal to %s.",
	},
	"instanceof": {
		Code:    "INVALID_INSTANCE",
		Message: "The value must be an instance of %s.",
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
