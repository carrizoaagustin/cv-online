package repository

import (
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

func (r *ResourceRepository) Create() {
	// USE EXAMPLE FOR MIGUEL
	r.queryBuilder.StartQuery().From("resources")
}
