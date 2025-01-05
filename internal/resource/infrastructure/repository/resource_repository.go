package repository

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"

	"github.com/carrizoaagustin/cv-online/internal/resource/domain"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/model"
	"github.com/carrizoaagustin/cv-online/pkg/dbquerybuilder"
)

type resourceDB struct {
	ResourceID uuid.UUID `db:"resource_id"`
	Filename   string
	Format     string
	Link       string
}

func transformToModel(dbObject resourceDB) *model.Resource {
	return &model.Resource{
		ID:       dbObject.ResourceID,
		Filename: dbObject.Filename,
		Format:   dbObject.Format,
		Link:     dbObject.Link,
	}
}

func transformToModelArray(dbObjects []resourceDB) []model.Resource {
	var resourcesList []model.Resource

	for _, element := range dbObjects {
		resourcesList = append(resourcesList, model.Resource{
			ID:       element.ResourceID,
			Filename: element.Filename,
			Format:   element.Format,
			Link:     element.Link,
		})
	}

	return resourcesList
}

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

func (r *ResourceRepositoryDB) Delete(resource model.Resource) error {
	_, err := r.queryBuilder.
		StartQuery().
		Delete("resources").
		Where(goqu.C("resource_id").Eq(resource.ID)).
		Executor().
		Exec()

	if err != nil {
		return err
	}

	return nil
}

func (r *ResourceRepositoryDB) GetByID(id uuid.UUID) (bool, *model.Resource, error) {
	var result resourceDB

	found, err := r.queryBuilder.StartQuery().
		From("resources").
		Where(goqu.C("resource_id").Eq(id)).
		ScanStruct(&result)

	if err != nil {
		return false, nil, err
	}

	return found, transformToModel(result), nil
}

func (r *ResourceRepositoryDB) FindResources() ([]model.Resource, error) {
	var results []resourceDB

	err := r.queryBuilder.StartQuery().
		From("resources").
		ScanStructs(&results)

	if err != nil {
		return nil, err
	}

	return transformToModelArray(results), nil
}
