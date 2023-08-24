package models

import (
	"log"
	"platform/internal/client"
)

func Migrate() {
	err := client.PostgreSqlClient.Migrator().AutoMigrate(
		&Product{},
		&Campaign{},
		&Order{},
	)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Migration successfully completed")
}
