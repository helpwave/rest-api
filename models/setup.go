package models

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rest-api/logging"
	"strings"
)

var database *gorm.DB = nil

// SetupDatabase tries to connect to the database and sets DB, else it exits the process
func SetupDatabase(host, user, password, databaseName, port string) {
	log.Info().
		Str("host", host).
		Str("user", user).
		Str("password", "<omitted>").
		Str("database", databaseName).
		Str("port", port).
		Msg("connecting to postgres...")

	postgresDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, databaseName, port,
	)

	log.Debug().
		Str("dsn", strings.Replace(postgresDSN, password, "<omitted>", -1)).
		Msg("dsn generated")

	config := gorm.Config{
		Logger: logging.GormLogger{},
	}

	db, dbErr := gorm.Open(postgres.Open(postgresDSN), &config)
	if dbErr != nil {
		log.Fatal().Err(dbErr).Msg("Could not connect to database: ")
	}
	log.Info().Msg("connected to db")

	database = db
}

func GetDB(logCtx context.Context) *gorm.DB {
	switch logCtx.(type) {
	case *gin.Context:
		log.Warn().Msg("logCtx is of type gin.Context. You probably passed the wrong context to GetDB()!")
	}
	if logCtx != nil {
		return database.WithContext(logCtx)
	}
	return database
}
