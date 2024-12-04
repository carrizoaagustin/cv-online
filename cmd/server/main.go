package main

import (
	"github.com/carrizoaagustin/cv-online/internal/resource/application/usecase"
	"github.com/carrizoaagustin/cv-online/internal/resource/infrastructure/storage"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/carrizoaagustin/cv-online/config"
	"github.com/carrizoaagustin/cv-online/pkg/dbconnection"
	"github.com/carrizoaagustin/cv-online/pkg/dbquerybuilder"

	resource_service "github.com/carrizoaagustin/cv-online/internal/resource/domain/service"
	resource_repository "github.com/carrizoaagustin/cv-online/internal/resource/infrastructure/repository"
	resource_controller "github.com/carrizoaagustin/cv-online/internal/resource/presentation/controller"
)

func main() {
	cfg := config.LoadConfig()
	databaseConnection := dbconnection.New(&cfg.DatabaseConfig)
	databaseConnection.CreateSchema()
	databaseConnection.Connect()
	defer databaseConnection.Close()

	databaseConnection.RunMigrations()

	// INIT OBJECTS
	queryBuilder := dbquerybuilder.New(databaseConnection.GetDatabaseConnection())

	// FILE STORAGE SERVICE

	// RESOURCES
	resourceRepository := resource_repository.NewResourceRepository(queryBuilder)

	resource_service.NewResourceService(resourceRepository)

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
