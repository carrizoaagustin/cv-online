package usecase

import (
	"github.com/carrizoaagustin/cv-online/config"
	"github.com/carrizoaagustin/cv-online/internal/resource/application"
	"github.com/carrizoaagustin/cv-online/internal/resource/application/dto"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/failures"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/service"
	"github.com/carrizoaagustin/cv-online/pkg/apperrors"
)

type ResourceUseCase struct {
	resourceService    service.ResourceService
	fileStorageService domain.FileStorageService
	config             config.ResourceConfig
}

func NewResourceUseCase(config config.ResourceConfig, fileStorageService domain.FileStorageService, resourceService service.ResourceService) application.ResourceUseCase {
	return &ResourceUseCase{
		config:             config,
		fileStorageService: fileStorageService,
		resourceService:    resourceService,
	}
}

func (u *ResourceUseCase) createLink(baseFolder string, baseUrl string, filename string) string {
	if baseFolder == "" {
		return baseUrl + "/" + filename
	}

	return baseUrl + "/" + baseFolder + "/" + filename
}

func (u *ResourceUseCase) UploadResource(input dto.UploadResourceDTO) error {
	link := u.createLink(u.config.BaseFolder, u.config.BaseURL, input.Filename)
	err := u.resourceService.Create(input.ConvertToCreateResourceData(link))

	if err != nil {
		return err
	}

	_, err = u.fileStorageService.UploadFile(input.ConvertToFileInput([]string{u.config.BaseFolder}))

	if err != nil {
		// TODO: Delete resource
		return apperrors.NewInternalError(failures.UploadError)

	}

	return nil
}
