package testutils

import (
	"context"
	"strings"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/carrizoaagustin/cv-online/config"
	"github.com/carrizoaagustin/cv-online/pkg/dbconnection"
	"github.com/carrizoaagustin/cv-online/pkg/dbquerybuilder"
)

type TestContainerPSQL struct {
	container *postgres.PostgresContainer
	ctx       context.Context
}

func StartPSQLContainer() *TestContainerPSQL {
	ctx := context.Background()

	postgresContainer, err := postgres.Run(ctx,
		"postgres:16",
		postgres.WithUsername("postgres"),
		postgres.WithPassword("root1234"),
		testcontainers.WithWaitStrategy(
			//nolint:mnd // It is for test porpuse
			wait.
				ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)

	if err != nil {
		panic(err)
	}

	return &TestContainerPSQL{
		container: postgresContainer,
		ctx:       ctx,
	}
}

func (t *TestContainerPSQL) Stop() {
	if err := t.container.Terminate(t.ctx); err != nil {
		panic(err)
	}
}

type TestSchema struct {
	queryBuilder *dbquerybuilder.DBQueryBuilder
	dbConnection *dbconnection.DatabaseConnection
}

func (t *TestSchema) CloseConnection() {
	t.dbConnection.Close()
}

func (t *TestSchema) GetQueryBuilder() *dbquerybuilder.DBQueryBuilder {
	return t.queryBuilder
}

func (t *TestContainerPSQL) StartTestSchema(schema string) *TestSchema {
	url, err := t.container.ConnectionString(t.ctx)

	if err != nil {
		panic(err)
	}

	lastSlashIndex := strings.LastIndex(url, "/")
	if lastSlashIndex != -1 {
		url = url[:lastSlashIndex]
	}

	databaseConnection := dbconnection.New(&config.DatabaseConfig{
		URL:        url,
		SchemaName: schema,
		SSLMode:    "disable",
	})

	databaseConnection.CreateSchema()
	databaseConnection.Connect()
	databaseConnection.RunMigrations()

	return &TestSchema{
		dbConnection: databaseConnection,
		queryBuilder: dbquerybuilder.New(databaseConnection.GetDatabaseConnection()),
	}
}
