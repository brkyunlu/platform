package client

import (
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"os"
	"platform/internal/database"
)

func GetPostgreSqlClient() *gorm.DB {
	var (
		postgresHost = os.Getenv("POSTGRES_HOST")
		postgresUser = os.Getenv("POSTGRES_USER")
		postgresPass = os.Getenv("POSTGRES_PASS")
		postgresPort = os.Getenv("POSTGRES_PORT")
		postgresDb   = os.Getenv("POSTGRES_DB")
		postgresSsl  = os.Getenv("POSTGRES_SSL")
	)

	return database.GetPostgreSqlConnection(postgresHost, postgresUser, postgresPass, postgresDb, postgresPort, postgresSsl)
}
