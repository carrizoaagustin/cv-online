package failures

import "github.com/carrizoaagustin/cv-online/pkg/apperrors"

//nolint:gochecknoglobals // I need global vars for errors
var (
	ResourceInvalidLinkError = apperrors.ErrorContent{
		Code:    "RESOURCE_INVALID_LINK",
		Message: "Invalid link",
	}
	ResourceInvalidFilenameError = apperrors.ErrorContent{
		Code:    "RESOURCE_INVALID_FILENAME",
		Message: "Invalid filename",
	}
	ResourceInvalidFormatError = apperrors.ErrorContent{
		Code:    "RESOURCE_INVALID_FORMAT",
		Message: "Invalid format. Only allow pdf, jpeg, png and gif formats",
	}
	ResourceCreationUnexpectedError = apperrors.ErrorContent{
		Code:    "RESOURCE_UNEXPECTED_ERROR",
		Message: "Unexpected error when create resource.",
	}
	UploadError = apperrors.ErrorContent{
		Code:    "RESOURCE_UPLOAD_ERROR",
		Message: "Error uploading file.",
	}
)
