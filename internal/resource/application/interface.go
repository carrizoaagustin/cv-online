package application

import "github.com/carrizoaagustin/cv-online/internal/resource/application/dto"

type ResourceUseCase interface {
	UploadResource(dto dto.UploadResourceDTO) error
}
