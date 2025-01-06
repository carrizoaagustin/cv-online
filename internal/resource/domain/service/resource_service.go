package service

import (
	"github.com/google/uuid"

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

func (s *ResourceService) Create(data dto.CreateResourceData) (*model.Resource, error) {
	resource, err := model.NewResource(data.Filename, data.Format, data.Link)
	if err != nil {
		return nil, err
	}

	err = s.resourceRepository.Create(*resource)
	if err != nil {
		return nil, apperrors.NewInternalError(failures.ResourceCreationError)
	}

	return resource, nil
}

func (s *ResourceService) Delete(id uuid.UUID) error {
	found, resource, err := s.resourceRepository.GetByID(id)

	if err != nil {
		return apperrors.NewInternalError(failures.ResourceGetError)
	}

	if found {
		err = s.resourceRepository.Delete(*resource)
		if err != nil {
			return apperrors.NewInternalError(failures.ResourceDeleteError)
		}
	}

	return nil
}

func (s *ResourceService) Find() ([]model.Resource, error) {
	resources, err := s.resourceRepository.FindResources()

	if err != nil {
		return nil, apperrors.NewInternalError(failures.ResourceFindError)
	}

	return resources, nil
}
