package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/carrizoaagustin/cv-online/config"
	"github.com/carrizoaagustin/cv-online/pkg/dbconnection"
)

func main() {
	cfg := config.LoadConfig()
	db := dbconnection.ConnectDB(&cfg.DatabaseConfig)
	defer db.Close()

	dbconnection.RunMigrations(db)

	if cfg.App.EnvironmentMode == config.ProductionMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	// DELETE THAT
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err := router.Run(":" + cfg.App.PORT)

	if err != nil {
		panic(err)
	}
}
