package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/carrizoaagustin/cv-online/internal/resource/domain/dto"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/failures"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/model"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/service"
	"github.com/carrizoaagustin/cv-online/pkg/apperrors"
)

type MockResourceRepository struct {
	mock.Mock
}

func (m *MockResourceRepository) Create(resource model.Resource) error {
	args := m.Called(resource)

	return args.Error(0)
}

func TestResourceService(t *testing.T) {
	type Given struct {
		createResourceData dto.CreateResourceData
		mockValue          error
	}

	type Expected struct {
		err error
	}

	testCases := map[string]struct {
		given    Given
		expected Expected
	}{
		"Creation success": {
			given: Given{
				createResourceData: dto.CreateResourceData{
					Link:     "https://test.image.com/image",
					Format:   model.Pdf,
					Filename: "test.pdf",
				},
				mockValue: nil,
			},
			expected: Expected{
				err: nil,
			},
		},
		"Invalid format": {
			given: Given{
				createResourceData: dto.CreateResourceData{
					Link:     "https://test.image.com/image",
					Format:   "invalid",
					Filename: "test.pdf",
				},
				mockValue: nil,
			},
			expected: Expected{
				err: apperrors.NewValidationError(failures.ResourceInvalidFormatError, "format"),
			},
		},
		"format empty": {
			given: Given{
				createResourceData: dto.CreateResourceData{
					Link:     "https://test.image.com/image",
					Format:   "",
					Filename: "test.pdf",
				},
				mockValue: nil,
			},
			expected: Expected{
				err: apperrors.NewValidationError(failures.ResourceInvalidFormatError, "format"),
			},
		},
		"Invalid link": {
			given: Given{
				createResourceData: dto.CreateResourceData{
					Link:     "",
					Format:   model.Pdf,
					Filename: "test.pdf",
				},
				mockValue: nil,
			},
			expected: Expected{
				err: apperrors.NewValidationError(failures.ResourceInvalidLinkError, "link"),
			},
		},
		"Invalid filename": {
			given: Given{
				createResourceData: dto.CreateResourceData{
					Link:     "https://test.image.com/image",
					Format:   model.Pdf,
					Filename: "",
				},
				mockValue: nil,
			},
			expected: Expected{
				err: apperrors.NewValidationError(failures.ResourceInvalidFilenameError, "filename"),
			},
		},
		"Multiple invalid fields": {
			given: Given{
				createResourceData: dto.CreateResourceData{
					Link:     "",
					Format:   "invalid format",
					Filename: "",
				},
				mockValue: nil,
			},
			expected: Expected{
				err: errors.Join(
					apperrors.NewValidationError(failures.ResourceInvalidFilenameError, "filename"),
					apperrors.NewValidationError(failures.ResourceInvalidFormatError, "format"),
					apperrors.NewValidationError(failures.ResourceInvalidLinkError, "link"),
				),
			},
		},
		"Repository error": {
			given: Given{
				createResourceData: dto.CreateResourceData{
					Link:     "https://test.image.com/image",
					Format:   model.Pdf,
					Filename: "filename",
				},
				mockValue: errors.New("test"),
			},
			expected: Expected{
				err: apperrors.NewInternalError(failures.ResourceCreationUnexpectedError),
			},
		},
	}

	for name, caseData := range testCases {
		t.Run(name, func(t *testing.T) {
			mockRepository := new(MockResourceRepository)
			resourceService := service.NewResourceService(mockRepository)
			mockRepository.
				On("Create", mock.Anything).
				Return(caseData.given.mockValue)

			_, err := resourceService.Create(caseData.given.createResourceData)

			if caseData.expected.err == nil {
				require.NoError(t, err, "Expected no error but got one")
			} else {
				require.Error(t, err)

				if joinedErr, ok := caseData.expected.err.(interface{ Unwrap() []error }); ok {
					for _, expectedErr := range joinedErr.Unwrap() {
						require.ErrorContains(t, err, expectedErr.Error(), "Expected error not found in error chain")
					}
				} else {
					require.EqualError(t, err, caseData.expected.err.Error(), "Error message does not match")
				}
			}
		})
	}
}
