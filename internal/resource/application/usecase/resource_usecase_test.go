package usecase_test

import (
	"crypto/rand"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/carrizoaagustin/cv-online/config"
	"github.com/carrizoaagustin/cv-online/internal/resource/application/dto"
	"github.com/carrizoaagustin/cv-online/internal/resource/application/usecase"
	dtodomain "github.com/carrizoaagustin/cv-online/internal/resource/domain/dto"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/failures"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/model"
	"github.com/carrizoaagustin/cv-online/pkg/apperrors"
)

type MockResourceService struct {
	mock.Mock
}

func (m *MockResourceService) Delete(id uuid.UUID) error {
	args := m.Called(id)

	return args.Error(0)
}

func (m *MockResourceService) Create(resource dtodomain.CreateResourceData) (*model.Resource, error) {
	args := m.Called(resource)

	return args.Get(0).(*model.Resource), args.Error(1)
}

type MockFileStorageService struct {
	mock.Mock
}

func (m *MockFileStorageService) UploadFile(file model.FileInput) (string, error) {
	args := m.Called(file)
	return args.String(0), args.Error(1)
}

func TestUploadResourceUseCase(t *testing.T) {
	type Given struct {
		input                       dto.UploadResourceDTO
		mockValueResourceService    error
		mockValueFileStorageService error
	}

	type Expected struct {
		err error
	}

	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		t.Fatal("Invalid byte generation")
	}

	testCases := map[string]struct {
		given    Given
		expected Expected
	}{
		"upload success": {
			given: Given{
				input: dto.UploadResourceDTO{
					File:        randomBytes,
					Filename:    "test.pdf",
					ContentType: model.Pdf,
				},
				mockValueResourceService:    nil,
				mockValueFileStorageService: nil,
			},
			expected: Expected{
				err: nil,
			},
		},
		"error with resource service": {
			given: Given{
				input: dto.UploadResourceDTO{
					File:        randomBytes,
					Filename:    "test.pdf",
					ContentType: model.Pdf,
				},
				mockValueResourceService:    errors.New("error in resource service"),
				mockValueFileStorageService: nil,
			},
			expected: Expected{
				err: errors.New("error in resource service"),
			},
		},
		"error with fileStorage service": {
			given: Given{
				input: dto.UploadResourceDTO{
					File:        randomBytes,
					Filename:    "test.pdf",
					ContentType: model.Pdf,
				},
				mockValueResourceService:    nil,
				mockValueFileStorageService: errors.New("error in file storage"),
			},
			expected: Expected{
				err: apperrors.NewInternalError(failures.UploadError),
			},
		},
	}

	for name, caseData := range testCases {
		t.Run(name, func(t *testing.T) {
			mockResourceService := new(MockResourceService)
			mockFileStorageService := new(MockFileStorageService)
			cfg := config.ResourceConfig{
				BaseFolder: "basefolder",
				BaseURL:    "https://test.image.com",
			}
			resourceUseCase := usecase.NewResourceUseCase(cfg, mockFileStorageService, mockResourceService)
			mockFileStorageService.
				On("UploadFile", mock.Anything).
				Return("https://test.image.com/resource", caseData.given.mockValueFileStorageService)

			mockResourceService.
				On("Create", mock.Anything).
				Return(&model.Resource{}, caseData.given.mockValueResourceService)

			mockResourceService.
				On("Delete", mock.Anything).
				Return(nil)

			err := resourceUseCase.UploadResource(caseData.given.input)
			if caseData.expected.err == nil {
				require.NoError(t, err, "Expected no error but got one")
			} else {
				require.EqualError(t, err, caseData.expected.err.Error(), "Error don't match")
			}
		})
	}
}
