package repository_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/carrizoaagustin/cv-online/internal/resource/domain"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/model"
	"github.com/carrizoaagustin/cv-online/internal/resource/infrastructure/repository"
	"github.com/carrizoaagustin/cv-online/testutils"
)

func TestInsertResource(t *testing.T) {
	type Given struct {
		resource model.Resource
	}

	type Expected struct {
		wantErr bool
	}

	tests := map[string]struct {
		given    Given
		expected Expected
		setup    func(resourceRepository domain.ResourceRepository)
	}{
		"Creation success": {
			given: Given{
				resource: model.Resource{ID: uuid.New(), Format: "pdf", Link: "https://asas", Filename: "filename"},
			},
			expected: Expected{
				wantErr: false,
			},
		},
		"query error": {
			given: Given{
				resource: model.Resource{ID: uuid.Nil, Format: "pdf", Link: "https://asas", Filename: "filename"},
			},
			expected: Expected{
				wantErr: true,
			},
			setup: func(resourceRepository domain.ResourceRepository) {
				err := resourceRepository.Create(model.Resource{ID: uuid.Nil, Format: "pdf", Link: "https://asas", Filename: "filename"})

				if err != nil {
					t.Errorf("Setup error. %v", err)
				}
			},
		},
	}

	psqlContainer := testutils.StartPSQLContainer()
	defer psqlContainer.Stop()

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			testSchema := psqlContainer.StartTestSchema(uuid.NewString())

			t.Cleanup(func() {
				testSchema.CloseConnection()
			})

			resourceRepository := repository.NewResourceRepository(testSchema.GetQueryBuilder())

			if test.setup != nil {
				test.setup(resourceRepository)
			}

			err := resourceRepository.Create(test.given.resource)

			if test.expected.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
