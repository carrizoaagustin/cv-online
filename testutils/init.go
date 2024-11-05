package testutils

import (
	"github.com/carrizoaagustin/cv-online/config"
	"github.com/carrizoaagustin/cv-online/pkg/dbconnection"
	"github.com/carrizoaagustin/cv-online/pkg/dbquerybuilder"
)

type TestDBResources struct {
	DBquerybuilder *dbquerybuilder.DBQueryBuilder
	DBconnection   *dbconnection.DatabaseConnection
}

func InitTestDBResources() *TestDBResources {
	cfg := config.LoadTestConfig()

	databaseConnection := dbconnection.New(&cfg.DatabaseConfig)
	databaseConnection.Connect()

	// INIT OBJECTS
	queryBuilder := dbquerybuilder.New(databaseConnection.GetDatabaseConnection())

	return &TestDBResources{
		DBquerybuilder: queryBuilder,
		DBconnection:   databaseConnection,
	}
}

func CloseTestDBResources(resource *TestDBResources) {
	resource.DBconnection.Close()
}
