package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"rest-api/logging"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log, _ := logging.GetRequestLogger(ctx)
		authHeader := ctx.GetHeader("Authorization")

		regex := regexp.MustCompile("^Bearer (\\w+)$")

		var err error
		var token string

		if len(authHeader) == 0 {
			err = errors.New("missing Authorization header")
		} else {
			matches := regex.FindStringSubmatch(authHeader)
			if len(matches) != 2 {
				err = errors.New("authorization header invalid")
			} else {
				token = matches[1]
			}
		}

		if err != nil {
			log.Warn().Err(err).Send()
			SendError(ctx, http.StatusUnauthorized, err)
			return
		}

		// TODO: validate token and set user
		ctx.Set("authToken", token)
	}
}
