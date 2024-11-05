package services

import (
	"github.com/carrizoaagustin/cv-online/internal/resource/domain"
)

type ResourceService struct {
	resourceRepository domain.ResourceRepository
}

func NewResourceService(repository domain.ResourceRepository) domain.ResourceService {
	return &ResourceService{
		resourceRepository: repository,
	}
}

func (s *ResourceService) Create(resource domain.Resource) error {
	s.resourceRepository.Create(resource)

	return nil
}
