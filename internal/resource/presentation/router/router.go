package router

import (
	"net/http"

	"github.com/carrizoaagustin/cv-online/auth"
	"github.com/carrizoaagustin/cv-online/auth/providers"
	"github.com/gin-gonic/gin"

	"github.com/carrizoaagustin/cv-online/internal/resource/presentation/controller"
)

func RegisterRoutes(router *gin.RouterGroup, controller *controller.ResourceController) {
	googleProvider := providers.NewGoogleProvider(
		auth.OAuthConfigParams{
			ClientID:  "240652747137-s3krn73b2qagrojnftun4cirpipr5pc2.apps.googleusercontent.com",
			SecretKey: "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",
			Scopes:    nil,
		},
	)
	contextGroup := router.Group("/resources")
	contextGroup.GET("", controller.Find)
	contextGroup.POST("", controller.UploadFile)
	contextGroup.GET("/login", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, googleProvider.GetAuthorizationURL())
	})
	contextGroup.GET("/callback", func(c *gin.Context) {
		code := c.Query("code")

		token, err := googleProvider.RedeemToken(code)
		if err != nil {
			c.Error(err)
			return
		}
		user, _ := googleProvider.GetUserInformation(token)
		if err != nil {
			c.Error(err)
			return
		}
		println(user.Email)

		c.JSON(200, user)
	})
}
