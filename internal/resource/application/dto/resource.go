package dto

import (
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/dto"
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
