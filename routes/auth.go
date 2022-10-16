package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/logging"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log, _ := logging.GetRequestLogger(ctx)
		authHeader := ctx.GetHeader("Authorization")

		if len(authHeader) == 0 {
			err := errors.New("missing Authorization header")
			log.Warn().Err(err).Send()
			SendError(ctx, http.StatusUnauthorized, err)
			return
		}
	}
}
