package repository_test

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/carrizoaagustin/cv-online/internal/resource/domain"
	"github.com/carrizoaagustin/cv-online/internal/resource/infrastructure/repository"
	"github.com/carrizoaagustin/cv-online/testutils"
)

//nolint:gochecknoglobals //i didn't found an alternative way
var dbResources *testutils.TestDBResources

func TestMain(m *testing.M) {
	dbResources = testutils.InitTestDBResources()

	// Ejecutar los tests
	code := m.Run()

	testutils.CloseTestDBResources(dbResources)

	os.Exit(code)
}

func TestInsertResource(t *testing.T) {
	type Given struct {
		resource domain.Resource
	}

	type Expected struct {
		wantErr bool
	}

	tests := map[string]struct {
		given    Given
		expected Expected
		setup    func(resourceRepository *repository.ResourceRepository)
	}{
		"Creation success": {
			given: Given{
				resource: domain.Resource{ID: uuid.New(), Format: "pdf", Link: "https://asas"},
			},
			expected: Expected{
				wantErr: false,
			},
		},
		"query error": {
			given: Given{
				resource: domain.Resource{ID: uuid.Nil, Format: "pdf", Link: "https://asas"},
			},
			expected: Expected{
				wantErr: true,
			},
			setup: func(resourceRepository *repository.ResourceRepository) {
				err := resourceRepository.Create(domain.Resource{ID: uuid.Nil, Format: "pdf", Link: "https://asas"})

				if err != nil {
					t.Errorf("Setup error. %v", err)
				}
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Cleanup(func() {
				_, err := dbResources.DBquerybuilder.StartQuery().Truncate("resources").Executor().Exec()

				if err != nil {
					t.Errorf("Error cleaning resources table")
				}
			})

			resourceRepository := repository.NewResourceRepository(dbResources.DBquerybuilder)

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