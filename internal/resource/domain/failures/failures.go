package failures

import "github.com/carrizoaagustin/cv-online/pkg/apperrors"

//nolint:gochecknoglobals // I need global vars for errors
var (
	ResourceInvalidLinkError = apperrors.ErrorContent{
		Message: "Invalid link",
		Code:    "RESOURCE_INVALID_LINK",
	}
	ResourceInvalidFormatError = apperrors.ErrorContent{
		Message: "Invalid format. Only allow pdf, jpeg, png and gif formats",
		Code:    "RESOURCE_INVALID_FORMAT",
	}
	ResourceCreationUnexpectedError = apperrors.ErrorContent{
		Message: "Unexpected error when create resource.",
		Code:    "RESOURCE_UNEXPECTED_ERROR",
	}
)
