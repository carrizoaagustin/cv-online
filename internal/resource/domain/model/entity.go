package model

import (
	"errors"
	"slices"

	"github.com/google/uuid"

	"github.com/carrizoaagustin/cv-online/internal/resource/domain/failures"
	"github.com/carrizoaagustin/cv-online/pkg/apperrors"
)

type Resource struct {
	ID       uuid.UUID
	Filename string
	Format   string
	Link     string
	UserID   uuid.UUID
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

func NewResource(filename string, format string, link string) (*Resource, error) {
	var errs []error

	if !isValidFormat(format) {
		errs = append(errs, apperrors.NewValidationError(failures.ResourceInvalidFormatError, "format"))
	}

	if link == "" {
		errs = append(errs, apperrors.NewValidationError(failures.ResourceInvalidLinkError, "link"))
	}

	if filename == "" {
		errs = append(errs, apperrors.NewValidationError(failures.ResourceInvalidFilenameError, "filename"))
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return &Resource{
		ID:       uuid.New(),
		Filename: filename,
		Format:   format,
		Link:     link,
	}, nil
}
