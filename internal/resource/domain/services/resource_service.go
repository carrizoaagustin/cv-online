package services

import (
	"github.com/carrizoaagustin/cv-online/internal/resource/domain"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/model"
)

type ResourceService struct {
	resourceRepository domain.ResourceRepository
}

func NewResourceService(repository domain.ResourceRepository) domain.ResourceService {
	return &ResourceService{
		resourceRepository: repository,
	}
}

func (s *ResourceService) Create(resource model.Resource) error {
	s.resourceRepository.Create(resource)

	return nil
}
