package domain

import (
	"github.com/google/uuid"

	"github.com/carrizoaagustin/cv-online/internal/resource/domain/dto"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/model"
)

type ResourceService interface {
	Create(resource dto.CreateResourceData) (*model.Resource, error)
	Delete(id uuid.UUID) error
	Find() ([]model.Resource, error)
}

type FileStorageService interface {
	UploadFile(file model.FileInput) (string, error)
}
