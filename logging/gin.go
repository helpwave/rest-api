package logging

import (
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

func GetRequestLogger(ctx *gin.Context) zerolog.Logger {
	return log.With().
		Str("request_id", requestid.Get(ctx)).
		Logger()
}
