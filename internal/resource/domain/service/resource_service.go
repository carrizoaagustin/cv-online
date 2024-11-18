package service

import (
	"github.com/carrizoaagustin/cv-online/internal/resource/domain"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/dto"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/failures"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/model"
	"github.com/carrizoaagustin/cv-online/pkg/apperrors"
)

type ResourceService struct {
	resourceRepository domain.ResourceRepository
}

func NewResourceService(repository domain.ResourceRepository) domain.ResourceService {
	return &ResourceService{
		resourceRepository: repository,
	}
}

func (s *ResourceService) Create(data dto.CreateResourceData) error {
	resource, err := model.NewResource(data.Format, data.Link)
	if err != nil {
		return err
	}

	err = s.resourceRepository.Create(*resource)
	if err != nil {
		return apperrors.NewInternalError(failures.ResourceCreationUnexpectedError)
	}

	return nil
}
