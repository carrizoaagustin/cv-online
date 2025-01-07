package application

import "github.com/carrizoaagustin/cv-online/internal/resource/application/dto"

type ResourceUseCase interface {
	Find() ([]dto.ResourceOutput, error)
	UploadResource(dto dto.UploadResourceDTO) error
}
