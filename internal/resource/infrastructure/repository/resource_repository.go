package repository

import (
	"github.com/carrizoaagustin/cv-online/internal/resource/domain"
	"github.com/carrizoaagustin/cv-online/pkg/dbquerybuilder"
)

type ResourceRepository struct {
	queryBuilder *dbquerybuilder.DBQueryBuilder
}

func NewResourceRepository(queryBuilder *dbquerybuilder.DBQueryBuilder) *ResourceRepository {
	return &ResourceRepository{
		queryBuilder: queryBuilder,
	}
}

func (r *ResourceRepository) Create(resource domain.Resource) error {
	// USE EXAMPLE FOR MIGUEL
	_, err := r.queryBuilder.StartQuery().Insert("resources").Rows(
		dbquerybuilder.Record{"resource_id": resource.ID, "format": resource.Format},
	).Executor().Exec()

	if err != nil {
		return err
	}

	return nil
}
