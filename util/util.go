package util

import (
	"fmt"
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
