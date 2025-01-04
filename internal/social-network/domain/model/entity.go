package model

import (
	"errors"

	"github.com/google/uuid"

	"github.com/carrizoaagustin/cv-online/internal/resource/domain/failures"
	"github.com/carrizoaagustin/cv-online/pkg/apperrors"
)

type SocialNetwork struct {
	ID   uuid.UUID
	Name string
}

func NewSocialNetwork(name string) (*SocialNetwork, error) {
	var errs []error

	if name == "" {
		errs = append(errs, apperrors.NewValidationError(failures.ResourceInvalidLinkError, "name"))
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return &SocialNetwork{
		ID:   uuid.New(),
		Name: name,
	}, nil

}
