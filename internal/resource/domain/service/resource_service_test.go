package service_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
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
func (m *MockResourceRepository) Delete(resource model.Resource) error {
	args := m.Called(resource)

	return args.Error(0)
}

func (m *MockResourceRepository) GetByID(id uuid.UUID) (bool, *model.Resource, error) {
	args := m.Called(id)
	var resource *model.Resource
	if args.Get(1) != nil {
		resource = args.Get(1).(*model.Resource)
	}
	return args.Bool(0), resource, args.Error(2)
}

func (m *MockResourceRepository) FindResources() ([]model.Resource, error) {
	args := m.Called()

	var resources []model.Resource
	if args.Get(0) != nil {
		resources = args.Get(0).([]model.Resource)
	}

	return resources, args.Error(1)
}

func TestCreateResourceService(t *testing.T) {
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
				err: apperrors.NewInternalError(failures.ResourceCreationError),
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

func TestFindResourceService(t *testing.T) {
	type Expected struct {
		err error
	}

	resource := model.Resource{
		ID:       uuid.New(),
		Filename: "Test",
		Format:   "pdf",
		Link:     "https://test.image.com/image",
	}

	testCases := map[string]struct {
		expected  Expected
		setupMock func() *MockResourceRepository
	}{
		"Find success": {
			expected: Expected{
				err: nil,
			},
			setupMock: func() *MockResourceRepository {
				mockRepository := new(MockResourceRepository)
				mockRepository.
					On("FindResources", mock.Anything).
					Return([]model.Resource{resource}, nil)
				return mockRepository
			},
		},
		"Find repository error": {
			expected: Expected{
				err: apperrors.NewInternalError(failures.ResourceFindError),
			},
			setupMock: func() *MockResourceRepository {
				mockRepository := new(MockResourceRepository)
				mockRepository.
					On("FindResources", mock.Anything).
					Return(nil, errors.New("test_error"))
				return mockRepository
			},
		},
	}

	for name, caseData := range testCases {
		t.Run(name, func(t *testing.T) {
			mockRepository := caseData.setupMock()
			resourceService := service.NewResourceService(mockRepository)

			_, err := resourceService.Find()

			if caseData.expected.err == nil {
				require.NoError(t, err, "Expected no error but got one")
			} else {
				require.Error(t, err)
				require.EqualError(t, err, caseData.expected.err.Error(), "Error message does not match")
			}
		})
	}
}

func TestDeleteResourceService(t *testing.T) {
	type Given struct {
		ID uuid.UUID
	}

	type Expected struct {
		err error
	}

	resource := model.Resource{
		ID:       uuid.New(),
		Filename: "Test",
		Format:   "pdf",
		Link:     "https://test.image.com/image",
	}

	testCases := map[string]struct {
		given     Given
		expected  Expected
		setupMock func() *MockResourceRepository
	}{
		"Delete success": {
			given: Given{
				ID: resource.ID,
			},
			expected: Expected{
				err: nil,
			},
			setupMock: func() *MockResourceRepository {
				mockRepository := new(MockResourceRepository)
				mockRepository.
					On("GetByID", mock.Anything).
					Return(true, &resource, nil)

				mockRepository.
					On("Delete", mock.Anything).
					Return(nil)

				return mockRepository
			},
		},
		"Delete nonexistent resource": {
			given: Given{
				ID: resource.ID,
			},
			expected: Expected{
				err: nil,
			},
			setupMock: func() *MockResourceRepository {
				mockRepository := new(MockResourceRepository)
				mockRepository.
					On("GetByID", mock.Anything).
					Return(false, nil, nil)

				mockRepository.
					On("Delete", mock.Anything).
					Return(nil)

				return mockRepository
			},
		},
		"Repository error - GetById": {
			given: Given{
				ID: resource.ID,
			},
			expected: Expected{
				err: apperrors.NewInternalError(failures.ResourceGetError),
			},
			setupMock: func() *MockResourceRepository {
				mockRepository := new(MockResourceRepository)
				mockRepository.
					On("GetByID", mock.Anything).
					Return(false, nil, errors.New("test_error"))

				mockRepository.
					On("Delete", mock.Anything).
					Return(nil)

				return mockRepository
			},
		},
		"Repository error - Delete": {
			given: Given{
				ID: resource.ID,
			},
			expected: Expected{
				err: apperrors.NewInternalError(failures.ResourceDeleteError),
			},
			setupMock: func() *MockResourceRepository {
				mockRepository := new(MockResourceRepository)
				mockRepository.
					On("GetByID", mock.Anything).
					Return(true, &resource, nil)

				mockRepository.
					On("Delete", mock.Anything).
					Return(errors.New("test_error"))

				return mockRepository
			},
		},
	}

	for name, caseData := range testCases {
		t.Run(name, func(t *testing.T) {
			mockRepository := caseData.setupMock()
			resourceService := service.NewResourceService(mockRepository)

			err := resourceService.Delete(caseData.given.ID)

			if caseData.expected.err == nil {
				require.NoError(t, err, "Expected no error but got one")
			} else {
				require.Error(t, err)
				require.EqualError(t, err, caseData.expected.err.Error(), "Error message does not match")
			}
		})
	}
}
