package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/carrizoaagustin/cv-online/config"
	resource_repository "github.com/carrizoaagustin/cv-online/internal/resource/infrastructure/repository"
	"github.com/carrizoaagustin/cv-online/pkg/dbconnection"
	"github.com/carrizoaagustin/cv-online/pkg/dbquerybuilder"
)

func main() {
	cfg := config.LoadConfig()
	databaseConnection := dbconnection.New(&cfg.DatabaseConfig)
	databaseConnection.Connect()
	defer databaseConnection.Close()
	

	databaseConnection.RunMigrations()

	// INIT OBJECTS
	queryBuilder := dbquerybuilder.New(databaseConnection.GetDatabaseConnection())
	resource_repository.NewResourceRepository(queryBuilder)

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
