package util

import (
	"fmt"
	"github.com/rs/zerolog/log"
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

// ReadFileOrEmpty will return the contents of a file or "" in case of error
func ReadFileOrEmpty(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Debug().Err(err).Msg("could not read " + path)
		return ""
	}
	return string(data)
}
