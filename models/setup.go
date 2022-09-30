package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB = nil

// SetupDatabase tries to connect to the database and sets DB, else it exits the process
func SetupDatabase(host, user, password, database, port string) {
	postgresDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, database, port,
	)

	db, dbErr := gorm.Open(postgres.Open(postgresDSN))
	if dbErr != nil {
		log.Fatalln("Could not connect to database: ", dbErr)
	}

	DB = db
}
