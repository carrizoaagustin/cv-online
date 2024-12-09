package repository

import (
	"github.com/carrizoaagustin/cv-online/internal/resource/domain"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/model"
	"github.com/carrizoaagustin/cv-online/pkg/dbquerybuilder"
)

type ResourceRepositoryDB struct {
	queryBuilder *dbquerybuilder.DBQueryBuilder
}

func NewResourceRepository(queryBuilder *dbquerybuilder.DBQueryBuilder) domain.ResourceRepository {
	return &ResourceRepositoryDB{
		queryBuilder: queryBuilder,
	}
}

func (r *ResourceRepositoryDB) Create(resource model.Resource) error {
	// USE EXAMPLE FOR MIGUEL
	_, err := r.queryBuilder.
		StartQuery().
		Insert("resources").
		Rows(dbquerybuilder.Record{"resource_id": resource.ID, "filename": resource.Filename, "format": resource.Format, "link": resource.Link}).
		Executor().
		Exec()

	if err != nil {
		return err
	}

	return nil
}
