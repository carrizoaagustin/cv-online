package router

import (
	"github.com/gin-gonic/gin"

	"github.com/carrizoaagustin/cv-online/internal/resource/presentation/controller"
)

func RegisterRoutes(router *gin.RouterGroup, controller *controller.ResourceController) {
	contextGroup := router.Group("/resources")

	contextGroup.GET("", controller.Find)
	contextGroup.POST("", controller.UploadFile)
}
