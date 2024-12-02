package controller

import (
	"io"

	"bytes"
	"github.com/gin-gonic/gin"

	"github.com/carrizoaagustin/cv-online/internal/resource/application"
	"github.com/carrizoaagustin/cv-online/internal/resource/application/dto"
)

type ResourceController struct {
	resourceUseCase application.ResourceUseCase
}

func NewResourceController(resourceUseCase application.ResourceUseCase) *ResourceController {
	return &ResourceController{
		resourceUseCase: resourceUseCase,
	}
}

func (c *ResourceController) UploadFile(ctx *gin.Context) {
	var form struct {
		Filename    string `form:"filename" binding:"required"`
		ContentType string `form:"content_type" binding:"required"`
	}

	err := ctx.ShouldBind(&form)
	if err != nil {
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		return
	}

	src, err := file.Open()
	if err != nil {
		return
	}

	defer src.Close()

	var buffer bytes.Buffer

	if _, err := io.Copy(&buffer, src); err != nil {
		return
	}

	input := dto.UploadResourceDTO{
		File:        buffer.Bytes(),
		Filename:    form.Filename,
		ContentType: form.ContentType,
	}

	err = c.resourceUseCase.UploadResource(input)
	if err != nil {
		return
	}

}
