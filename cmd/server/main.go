package main

import (
	"github.com/carrizoaagustin/cv-online/config"
	"github.com/carrizoaagustin/cv-online/pkg/dbconnection"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg :=config.LoadConfig()
	db := dbconnection.ConnectDB(&cfg.DatabaseConfig)
	defer db.Close()

	dbconnection.RunMigrations(db)

	if(cfg.App.EnvironmentMode == config.ProductionMode){
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	// DELETE THAT
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":" + cfg.App.PORT)
}