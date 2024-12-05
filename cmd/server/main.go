package main

import (
	"github.com/carrizoaagustin/cv-online/pkg/middleware"
	"github.com/gin-gonic/gin"

	"github.com/carrizoaagustin/cv-online/config"
	"github.com/carrizoaagustin/cv-online/pkg/dbconnection"
	"github.com/carrizoaagustin/cv-online/pkg/dbquerybuilder"

	resource_usecase "github.com/carrizoaagustin/cv-online/internal/resource/application/usecase"
	resource_service "github.com/carrizoaagustin/cv-online/internal/resource/domain/service"
	resource_repository "github.com/carrizoaagustin/cv-online/internal/resource/infrastructure/repository"
	resource_storage "github.com/carrizoaagustin/cv-online/internal/resource/infrastructure/storage"
	resource_controller "github.com/carrizoaagustin/cv-online/internal/resource/presentation/controller"
	resource_router "github.com/carrizoaagustin/cv-online/internal/resource/presentation/router"
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
	clientR2 := resource_storage.NewR2Client(cfg.StorageR2)
	fileStorageService := resource_storage.NewFileStorageServiceR2(cfg.StorageR2, clientR2)

	// RESOURCES
	resourceRepository := resource_repository.NewResourceRepository(queryBuilder)

	resourceService := resource_service.NewResourceService(resourceRepository)
	resourceUseCase := resource_usecase.NewResourceUseCase(cfg.ResourceConfig, fileStorageService, resourceService)
	resourceController := resource_controller.NewResourceController(resourceUseCase)

	// INIT ROUTER

	if cfg.App.EnvironmentMode == config.ProductionMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(middleware.ErrorHandler())
	api := router.Group("/api")

	// REGISTER ROUTES
	resource_router.RegisterRoutes(api, resourceController)

	err := router.Run(":" + cfg.App.PORT)

	if err != nil {
		panic(err)
	}
}
