package domain

import (
	"github.com/google/uuid"

	"github.com/carrizoaagustin/cv-online/internal/resource/domain/model"
)

type ResourceRepository interface {
	Create(resource model.Resource) error
	FindResources() ([]model.Resource, error)
	Delete(resource model.Resource) error
	GetByID(id uuid.UUID) (bool, *model.Resource, error)
}
