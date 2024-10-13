package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/carrizoaagustin/cv-online/config"
	"github.com/carrizoaagustin/cv-online/pkg/dbconnection"
)

func main() {
	cfg := config.LoadConfig()
	databaseConnection := dbconnection.New(&cfg.DatabaseConfig)
	databaseConnection.Connect()
	defer databaseConnection.Close()

	databaseConnection.RunMigrations()

	// INIT OBJECTS
	// queryBuilder := dbquerybuilder.New(databaseConnection.GetDatabaseConnection())

	// INIT ROUTER

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
