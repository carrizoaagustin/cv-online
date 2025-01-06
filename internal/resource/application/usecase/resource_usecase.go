package usecase

import (
	"github.com/google/uuid"

	"github.com/carrizoaagustin/cv-online/config"
	"github.com/carrizoaagustin/cv-online/internal/resource/application"
	"github.com/carrizoaagustin/cv-online/internal/resource/application/dto"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/failures"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/model"
	"github.com/carrizoaagustin/cv-online/pkg/apperrors"
)

type ResourceUseCase struct {
	resourceService    domain.ResourceService
	fileStorageService domain.FileStorageService
	config             config.ResourceConfig
}

func NewResourceUseCase(config config.ResourceConfig, fileStorageService domain.FileStorageService, resourceService domain.ResourceService) application.ResourceUseCase { //nolint:lll
	return &ResourceUseCase{
		config:             config,
		fileStorageService: fileStorageService,
		resourceService:    resourceService,
	}
}

func (u *ResourceUseCase) createLink(baseFolder string, baseURL string, filename string) string {
	if baseFolder == "" {
		return baseURL + "/" + filename
	}

	return baseURL + "/" + baseFolder + "/" + filename
}

func (u *ResourceUseCase) UploadResource(input dto.UploadResourceDTO) error {
	filename := uuid.New().String()
	link := u.createLink(u.config.BaseFolder, u.config.BaseURL, filename)
	resource, err := u.resourceService.Create(input.ConvertToCreateResourceData(link))

	if err != nil {
		return err
	}

	fileInput := model.FileInput{
		File:        input.File,
		Filename:    filename,
		ContentType: input.ContentType,
		Folders:     []string{u.config.BaseFolder},
	}

	_, err = u.fileStorageService.UploadFile(fileInput)
	if err != nil {
		_ = u.resourceService.Delete(resource.ID)
		return apperrors.NewInternalError(failures.UploadError)
	}
	return nil
}
