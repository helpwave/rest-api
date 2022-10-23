package logging

import (
	"context"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

func GinLogger(ctx *gin.Context, _ zerolog.Logger) zerolog.Logger {
	l := log.Logger

	// log otel context in case we ever use it
	otelContext := trace.SpanFromContext(ctx.Request.Context()).SpanContext()
	if otelContext.IsValid() {
		l.With().
			Str("trace_id", otelContext.TraceID().String()).
			Str("span_id", otelContext.SpanID().String()).
			Logger()
	}

	return l.With().
		Str("request_id", requestid.Get(ctx)).
		Logger()
}

// GetRequestLogger returns a logger with request information and a context
// You can use it like this:
//
//	func Handler(ctx *gin.Context) {
//		log, logCtx := logging.GetRequestLogger(ctx)
//		db := models.GetDB(logCtx)
//		log.Info().Msg("I'm a Request Handler!") // <- this message will contain request information
//		db.First(&...) // <- gorm logger contains request information
//		...
//	}
func GetRequestLogger(ctx *gin.Context) (zerolog.Logger, context.Context) {
	logger := log.With().
		Str("request_id", requestid.Get(ctx)).
		Logger()
	logCtx := logger.WithContext(ctx)
	return logger, logCtx
}
