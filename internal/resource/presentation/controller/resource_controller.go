package controller

import (
	"bytes"
	"io"

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

func (c *ResourceController) UploadFile(ctx *gin.Context) { // coverage-ignore
	var form struct {
		Filename    string `form:"filename" binding:"required"`
		ContentType string `form:"content_type" binding:"required"`
	}

	err := ctx.Bind(&form)
	if err != nil {
		ctx.Error(err)
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.Error(err)
		return
	}

	src, err := file.Open()
	if err != nil {
		ctx.Error(err)
		return
	}

	defer src.Close()

	var buffer bytes.Buffer

	if _, err = io.Copy(&buffer, src); err != nil {
		ctx.Error(err)
		return
	}

	input := dto.UploadResourceDTO{
		File:        buffer.Bytes(),
		Filename:    form.Filename,
		ContentType: form.ContentType,
	}

	err = c.resourceUseCase.UploadResource(input)
	if err != nil {
		ctx.Error(err)
		return
	}
}
