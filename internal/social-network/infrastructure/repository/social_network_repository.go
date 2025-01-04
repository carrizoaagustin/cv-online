package repository

import (
	"github.com/carrizoaagustin/cv-online/internal/social-network/domain"
	"github.com/carrizoaagustin/cv-online/internal/social-network/domain/model"
	"github.com/carrizoaagustin/cv-online/pkg/dbquerybuilder"
)

type SocialNetworkRepositoryDB struct {
	queryBuilder *dbquerybuilder.DBQueryBuilder
}

func NewSocialNetworkRepository(queryBuilder *dbquerybuilder.DBQueryBuilder) domain.SocialNetworkRepository {
	return &SocialNetworkRepositoryDB{
		queryBuilder: queryBuilder,
	}
}

func (r *SocialNetworkRepositoryDB) Create(socialNetwork model.SocialNetwork) error {
	// USE EXAMPLE FOR MIGUEL
	_, err := r.queryBuilder.
		StartQuery().
		Insert("social_networks").
		Rows(dbquerybuilder.Record{"social_network_id": socialNetwork.ID, "name": socialNetwork.Name}).
		Executor().
		Exec()

	if err != nil {
		return err
	}
	return nil
}
