package models

import (
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB = nil

// SetupDatabase tries to connect to the database and sets DB, else it exits the process
func SetupDatabase(host, user, password, database, port string) {
	postgresDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, database, port,
	)

	db, dbErr := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if dbErr != nil {
		log.Fatalln("Could not connect to database: ", dbErr)
	}

	DB = db
}

// IsOurFault returns true when a gORM error is caused by invalid backend configuration
// e.g. a disconnected database
func IsOurFault(gormError error) bool {
	// try to cast into PgError
	var pgErr *pgconn.PgError
	if errors.As(gormError, &pgErr) {
		return isOurFaultPgError(pgErr)
	}

	ours :=
		errors.Is(gormError, gorm.ErrInvalidTransaction) ||
			errors.Is(gormError, gorm.ErrNotImplemented) ||
			errors.Is(gormError, gorm.ErrMissingWhereClause) ||
			errors.Is(gormError, gorm.ErrUnsupportedRelation) ||
			errors.Is(gormError, gorm.ErrPrimaryKeyRequired) ||
			errors.Is(gormError, gorm.ErrModelValueRequired) ||
			errors.Is(gormError, gorm.ErrUnsupportedDriver) ||
			errors.Is(gormError, gorm.ErrRegistered) ||
			errors.Is(gormError, gorm.ErrInvalidField) ||
			errors.Is(gormError, gorm.ErrDryRunModeUnsupported) ||
			errors.Is(gormError, gorm.ErrInvalidDB) ||
			errors.Is(gormError, gorm.ErrInvalidValueOfLength) ||
			errors.Is(gormError, gorm.ErrPreloadNotAllowed)
	theirs :=
		errors.Is(gormError, gorm.ErrRecordNotFound) ||
			errors.Is(gormError, gorm.ErrInvalidData) ||
			errors.Is(gormError, gorm.ErrEmptySlice) ||
			errors.Is(gormError, gorm.ErrInvalidValue) ||
			errors.Is(gormError, gorm.ErrInvalidValueOfLength)

	if !ours && !theirs {
		log.Println("Error is neither theirs nor ours!!")
		return true // if we don't know the error, we might have caused it
	}

	return ours
}

func isOurFaultPgError(err *pgconn.PgError) bool {
	code := err.Code

	// common errors that can occur because the backend messed up
	return pgerrcode.IsConnectionException(code) ||
		pgerrcode.IsFeatureNotSupported(code) ||
		pgerrcode.IsInvalidTransactionInitiation(code) ||
		pgerrcode.IsInvalidSQLStatementName(code) ||
		pgerrcode.IsSyntaxErrororAccessRuleViolation(code) ||
		pgerrcode.IsInsufficientResources(code) ||
		pgerrcode.IsProgramLimitExceeded(code) ||
		pgerrcode.IsSystemError(code) ||
		pgerrcode.IsConfigurationFileError(code) ||
		pgerrcode.IsInternalError(code) ||
		pgerrcode.IsOperatorIntervention(code)
}
