package failures

import "github.com/carrizoaagustin/cv-online/pkg/apperrors"

//nolint:gochecknoglobals // I need global vars for errors
var (
	SocialNetworkInvalidNameError = apperrors.ErrorContent{
		Code:    "SOCIAL_NETWORK_INVALID_NAME",
		Message: "Invalid name",
	}
)
