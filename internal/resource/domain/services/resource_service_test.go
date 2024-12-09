package services_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/carrizoaagustin/cv-online/internal/resource/domain/dto"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/failures"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/model"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/services"
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
					Link:   "https://test.image.com/image",
					Format: model.Pdf,
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
					Link:   "https://test.image.com/image",
					Format: "invalid",
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
					Link:   "https://test.image.com/image",
					Format: "",
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
					Link:   "",
					Format: model.Pdf,
				},
				mockValue: nil,
			},
			expected: Expected{
				err: apperrors.NewValidationError(failures.ResourceInvalidLinkError, "link"),
			},
		},
		"Repository error": {
			given: Given{
				createResourceData: dto.CreateResourceData{
					Link:   "https://test.image.com/image",
					Format: model.Pdf,
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
			resourceService := services.NewResourceService(mockRepository)
			mockRepository.
				On("Create", mock.Anything).
				Return(caseData.given.mockValue)

			err := resourceService.Create(caseData.given.createResourceData)

			if caseData.expected.err == nil {
				// Caso de éxito: verificar que no haya error
				require.NoError(t, err, "Expected no error but got one")
			} else {
				// Caso de fallo: verificar el error esperado
				require.IsType(t, err, caseData.expected.err, "Error type don't match")
				require.EqualError(t, err, caseData.expected.err.Error(), "Error don't match")
			}
		})
	}
}
