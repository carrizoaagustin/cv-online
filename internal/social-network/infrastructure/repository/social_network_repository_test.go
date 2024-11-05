package repository_test

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/carrizoaagustin/cv-online/internal/social-network/domain"
	"github.com/carrizoaagustin/cv-online/internal/social-network/infrastructure/repository"
	"github.com/carrizoaagustin/cv-online/testutils"
)

//nolint:gochecknoglobals //i didn't found an alternative way
var dbResources *testutils.TestDBResources

func TestMain(m *testing.M) {
	//dbResources = testutils.InitTestDBResources()

	// Ejecutar los tests
	code := m.Run()

	testutils.CloseTestDBResources(dbResources)

	os.Exit(code)
}

func TestInsertSocialNetwork(t *testing.T) {
	type Given struct {
		socialNetwork domain.SocialNetwork
	}

	type Expected struct {
		wantErr bool
	}

	tests := map[string]struct {
		given    Given
		expected Expected
		setup    func(socialNetworkRepository domain.SocialNetworkRepository)
	}{
		"Creation success": {
			given: Given{
				socialNetwork: domain.SocialNetwork{ID: uuid.New(), Name: "github"},
			},
			expected: Expected{
				wantErr: false,
			},
		},
		"Query error": {
			given: Given{
				socialNetwork: domain.SocialNetwork{ID: uuid.Nil, Name: "github"},
			},
			expected: Expected{
				wantErr: true,
			},
			setup: func(socialNetworkRepository domain.SocialNetworkRepository) {
				err := socialNetworkRepository.Create(domain.SocialNetwork{ID: uuid.Nil, Name: "github"})

				if err != nil {
					t.Errorf("Setup error. %v", err)
				}
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Cleanup(func() {
				_, err := dbResources.DBquerybuilder.StartQuery().Truncate("social_networks").Executor().Exec()
				if err != nil {
					t.Errorf("Error cleaning social_networks table")
				}
			})

			socialNetworkRepository := repository.NewSocialNetworkRepository(dbResources.DBquerybuilder)

			if test.setup != nil {
				test.setup(socialNetworkRepository)
			}

			err := socialNetworkRepository.Create(test.given.socialNetwork)

			if test.expected.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
