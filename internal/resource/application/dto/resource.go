package dto

import (
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/dto"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain/model"
)

type UploadResourceDTO struct {
	File        []byte
	Filename    string
	ContentType string
}

func (r *UploadResourceDTO) ConvertToCreateResourceData(baseLink string) dto.CreateResourceData {
	return dto.CreateResourceData{
		Link:   baseLink + "/" + r.Filename,
		Format: r.ContentType,
	}
}

func (r *UploadResourceDTO) ConvertToFileInput(folders []string) model.FileInput {
	return model.FileInput{
		File:        r.File,
		Filename:    r.Filename,
		Folders:     folders,
		ContentType: r.ContentType,
	}

}
