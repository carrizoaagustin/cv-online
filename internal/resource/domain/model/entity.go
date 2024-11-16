package model

import (
	"slices"

	"github.com/google/uuid"

	"github.com/carrizoaagustin/cv-online/internal/resource/domain/failures"
	"github.com/carrizoaagustin/cv-online/pkg/apperrors"
)

type Resource struct {
	ID     uuid.UUID
	Format string
	Link   string
	UserID uuid.UUID
}

const (
	Pdf  = "application/pdf"
	Png  = "image/png"
	Jpeg = "image/jpeg"
	Gift = "image/gif"
)

func isValidFormat(format string) bool {
	allowedFormats := []string{Pdf, Png, Jpeg, Gift}
	return slices.Contains(allowedFormats, format)
}

func NewResource(format string, link string) (*Resource, error) {
	if !isValidFormat(format) {
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
