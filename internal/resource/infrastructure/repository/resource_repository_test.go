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

//nolint:gochecknoglobals // Because I need share the same container
var psqlContainer *testutils.TestContainerPSQL

func TestMain(m *testing.M) {
	psqlContainer = testutils.StartPSQLContainer()
	defer psqlContainer.Stop()
	m.Run()
}

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
		"Insert resource - creation success": {
			given: Given{
				resource: model.Resource{ID: uuid.New(), Format: "pdf", Link: "https://asas", Filename: "filename"},
			},
			expected: Expected{
				wantErr: false,
			},
		},
		"Insert resource - DB error": {
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

func TestDeleteResource(t *testing.T) {
	type Given struct {
		resource *model.Resource
	}

	type Expected struct {
		wantErr bool
	}

	resource := &model.Resource{ID: uuid.New(), Format: "pdf", Link: "https://asas", Filename: "filename"}
	totalResources := 5

	tests := map[string]struct {
		given               Given
		expected            Expected
		setup               func(resourceRepository domain.ResourceRepository)
		wantCloseConnection bool
	}{
		"Delete resource - success": {
			given: Given{
				resource,
			},
			expected: Expected{
				wantErr: false,
			},
			setup: func(resourceRepository domain.ResourceRepository) {
				err := resourceRepository.Create(*resource)

				if err != nil {
					t.Errorf("Error creating resource. %v", err)
				}

				for range totalResources {
					if err = resourceRepository.Create(model.Resource{ID: uuid.New(), Format: "pdf", Link: "https://asas", Filename: "filename"}); err != nil {
						t.Errorf("Error creating resource. %v", err)
					}
				}
			},
		},
		"Delete resource - non existent item": {
			given: Given{
				resource: &model.Resource{ID: uuid.Nil, Format: "pdf", Link: "https://asas", Filename: "filename"},
			},
			expected: Expected{
				wantErr: false,
			},
			setup: func(resourceRepository domain.ResourceRepository) {
				for range totalResources {
					if err := resourceRepository.Create(model.Resource{ID: uuid.New(), Format: "pdf", Link: "https://asas", Filename: "filename"}); err != nil {
						t.Errorf("Error creating resource. %v", err)
					}
				}
			},
		},
		"Delete resource - DB error": {
			given: Given{
				resource,
			},
			expected: Expected{
				wantErr: true,
			},
			wantCloseConnection: true,
		},
	}

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

			if test.wantCloseConnection {
				testSchema.CloseConnection()
			}

			err := resourceRepository.Delete(*test.given.resource)

			if test.expected.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				resourceList, _ := resourceRepository.FindResources()
				require.Len(t, resourceList, totalResources)
				found, _, _ := resourceRepository.GetByID(test.given.resource.ID)
				require.False(t, found, "Found object with ID when it should not be present in the DB")
			}
		})
	}
}

func TestFindResources(t *testing.T) {

	type Expected struct {
		totalResources int
		wantErr        bool
	}

	tests := map[string]struct {
		expected            Expected
		setup               func(resourceRepository domain.ResourceRepository)
		wantCloseConnection bool
	}{
		"Find resources - existing items": {
			expected: Expected{
				totalResources: 5,
			},
			setup: func(resourceRepository domain.ResourceRepository) {
				for range 5 {
					if err := resourceRepository.Create(model.Resource{ID: uuid.New(), Format: "pdf", Link: "https://asas", Filename: "filename"}); err != nil {
						t.Errorf("Error creating resource. %v", err)
					}
				}
			},
		},
		"Find resources - not existing items": {
			expected: Expected{
				totalResources: 0,
			},
		},
		"Find resources - DB error": {
			wantCloseConnection: true,
			expected: Expected{
				wantErr: true,
			},
		},
	}

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

			if test.wantCloseConnection {
				testSchema.CloseConnection()
			}

			items, err := resourceRepository.FindResources()

			if test.expected.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Len(t, items, test.expected.totalResources)
			}
		})
	}
}

func TestGetResource(t *testing.T) {
	type Given struct {
		resourceID uuid.UUID
	}

	type Expected struct {
		resourceExpected *model.Resource
		found            bool
		wantErr          bool
	}

	resource := &model.Resource{ID: uuid.New(), Format: "pdf", Link: "https://asas", Filename: "filename"}

	tests := map[string]struct {
		given               Given
		expected            Expected
		setup               func(resourceRepository domain.ResourceRepository)
		wantCloseConnection bool
	}{
		"Get success - existing item": {
			given: Given{
				resourceID: resource.ID,
			},
			expected: Expected{
				resourceExpected: resource,
				found:            true,
			},
			setup: func(resourceRepository domain.ResourceRepository) {
				err := resourceRepository.Create(*resource)

				if err != nil {
					t.Errorf("Error creating resource. %v", err)
				}
			},
		},
		"Get success - not existing item": {
			given: Given{
				resourceID: uuid.New(),
			},
			expected: Expected{
				resourceExpected: nil,
				found:            false,
			},
		},
		"Get resource - DB error": {
			wantCloseConnection: true,
			expected: Expected{
				wantErr: true,
			},
		},
	}

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

			if test.wantCloseConnection {
				testSchema.CloseConnection()
			}

			found, item, err := resourceRepository.GetByID(test.given.resourceID)

			if test.expected.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, test.expected.found, found)

				if test.expected.found {
					require.Equal(t, test.expected.resourceExpected.ID, item.ID)
					require.Equal(t, test.expected.resourceExpected.Filename, item.Filename)
					require.Equal(t, test.expected.resourceExpected.Format, item.Format)
					require.Equal(t, test.expected.resourceExpected.Link, item.Link)
				}
			}
		})
	}
}
