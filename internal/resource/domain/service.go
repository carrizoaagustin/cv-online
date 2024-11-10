package domain

import "github.com/carrizoaagustin/cv-online/internal/resource/domain/model"

type ResourceService interface {
	Create(resource model.Resource) error
}

type FileStorageService interface {
	UploadFile(file model.FileInput) (string, error)
}
