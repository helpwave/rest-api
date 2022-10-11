package logging

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func SetupLogging(mode, rawLevel string) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if mode != "release" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
	log.Logger = log.With().Caller().Logger()
	level, err := zerolog.ParseLevel(rawLevel)
	if err != nil {
		log.Fatal().Err(err).Msg("could not parse LOG_LEVEL")
	}
	log.Logger = log.Level(level)
	log.Info().Msg("Logging is set up")
}