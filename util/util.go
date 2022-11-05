package util

import (
	"fmt"
	"gorm.io/gorm"
	"os"
	"strings"
)

// GetEnvOr returns the environment variable named `key` or returns a default value
func GetEnvOr(key, fallback string) string {
	value, found := os.LookupEnv(key)
	if found {
		return value
	}
	return fallback
}

// Formatted formats anything but makes sure to encode newlines
func Formatted(arg any) string {
	return strings.Replace(fmt.Sprintf("%v", arg), "\n", "\\n", -1)
}

// GetTableName returns the table name of a model
// e.g.: `GetTableName(db, EmergencyRoom{})` will return "emergency_rooms"
func GetTableName(db *gorm.DB, model interface{}) string {
	stmt := &gorm.Statement{DB: db}
	_ = stmt.Parse(&model)
	return stmt.Schema.Table
}
