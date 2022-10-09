package models

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

var DB *gorm.DB = nil

// SetupDatabase tries to connect to the database and sets DB, else it exits the process
func SetupDatabase(host, user, password, database, port string) {
	log.Info().
		Str("host", host).
		Str("user", user).
		Str("password", "<omitted>").
		Str("database", database).
		Str("port", port).
		Msg("connecting to postgres...")

	postgresDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, database, port,
	)

	log.Debug().
		Str("dsn", strings.Replace(postgresDSN, password, "<omitted>", -1)).
		Msg("dsn generated")

	db, dbErr := gorm.Open(postgres.Open(postgresDSN))
	if dbErr != nil {
		log.Fatal().Err(dbErr).Msg("Could not connect to database: ")
	}
	log.Info().Msg("connected to db")

	DB = db
}
