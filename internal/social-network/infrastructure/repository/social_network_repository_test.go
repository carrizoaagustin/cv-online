package repository_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/carrizoaagustin/cv-online/internal/social-network/domain"
	"github.com/carrizoaagustin/cv-online/internal/social-network/domain/model"
	"github.com/carrizoaagustin/cv-online/internal/social-network/infrastructure/repository"
	"github.com/carrizoaagustin/cv-online/testutils"
)

func TestInsertSocialNetwork(t *testing.T) {
	type Given struct {
		socialNetwork model.SocialNetwork
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
				socialNetwork: model.SocialNetwork{ID: uuid.New(), Name: "github"},
			},
			expected: Expected{
				wantErr: false,
			},
		},
		"Query error": {
			given: Given{
				socialNetwork: model.SocialNetwork{ID: uuid.Nil, Name: "github"},
			},
			expected: Expected{
				wantErr: true,
			},
			setup: func(socialNetworkRepository domain.SocialNetworkRepository) {
				err := socialNetworkRepository.Create(model.SocialNetwork{ID: uuid.Nil, Name: "github"})

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

			socialNetworkRepository := repository.NewSocialNetworkRepository(testSchema.GetQueryBuilder())

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
