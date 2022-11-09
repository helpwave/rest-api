package routes

import (
	"context"
	"net/http"
	"rest-api/models"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type HTTPError struct {
	Code    int
	Message string
}

type HTTPErrorResponse struct {
	Error HTTPError
}

func SendError(ctx *gin.Context, status int, err error) {
	httpError := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	resp := HTTPErrorResponse{Error: httpError}
	ctx.JSON(status, resp)
	ctx.Abort()
}

func HandleDBError(ctx *gin.Context, logCtx context.Context, err error) {
	log.Ctx(logCtx).Warn().Err(err).Msg("db error occurred")
	status := http.StatusBadRequest
	if models.IsOurFault(err) {
		status = http.StatusInternalServerError
	}
	SendError(ctx, status, err)
}
