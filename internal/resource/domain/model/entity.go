package model

import (
	"slices"

	"github.com/carrizoaagustin/cv-online/internal/resource/domain/failures"
	"github.com/carrizoaagustin/cv-online/pkg/apperrors"
	"github.com/google/uuid"
)

type Resource struct {
	ID     uuid.UUID
	Format string
	Link   string
	UserID uuid.UUID
}

const (
	pdf  = "application/pdf"
	png  = "image/png"
	jpeg = "image/jpeg"
	gift = "image/gif"
)

func isValidFormat(format string) bool {
	allowedFormats := []string{pdf, png, jpeg, gift}
	return slices.Contains(allowedFormats, format)
}

func NewResource(format string, link string) (*Resource, error) {
	if isValidFormat(format) {
		return nil, apperrors.NewValidationError(failures.ResourceInvalidFormatError, "format")
	}

	if link == "" {
		return nil, apperrors.NewValidationError(failures.ResourceInvalidLinkError, "link")
	}

	return &Resource{
		ID:     uuid.New(),
		Format: format,
		Link:   link,
	}, nil
}
