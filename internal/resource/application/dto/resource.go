package dto

import (
	"github.com/google/uuid"

	"github.com/carrizoaagustin/cv-online/internal/resource/domain/dto"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/model"
)

type UploadResourceDTO struct {
	File        []byte
	Filename    string
	ContentType string
}

func (r *UploadResourceDTO) ConvertToCreateResourceData(link string) dto.CreateResourceData {
	return dto.CreateResourceData{
		Filename: r.Filename,
		Format:   r.ContentType,
		Link:     link,
	}
}

type ResourceOutput struct {
	ID       uuid.UUID `json:"id"`
	Filename string    `json:"filename"`
	Format   string    `json:"format"`
	Link     string    `json:"link"`
}

func TransformToResourceOutput(resource *model.Resource) *ResourceOutput {
	return &ResourceOutput{
		ID:       resource.ID,
		Filename: resource.Filename,
		Format:   resource.Format,
		Link:     resource.Link,
	}
}

func TransformToResourceOutputArray(resources []model.Resource) []ResourceOutput {
	var resourcesOutput []ResourceOutput

	for _, element := range resources {
		resourcesOutput = append(resourcesOutput, *TransformToResourceOutput(&element))
	}

	return resourcesOutput
}
