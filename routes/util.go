package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"rest-api/models"
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
}

// StatusResponse should give the caller feedback on the state of the task they submitted
type StatusResponse struct {
	Ok bool
}

// SendOk sends a basic "OK" response
func SendOk(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, StatusResponse{Ok: true})
}

func HandleDBError(ctx *gin.Context, logCtx context.Context, err error) {
	log.Ctx(logCtx).Warn().Err(err).Msg("db error occurred")
	status := http.StatusBadRequest
	if models.IsOurFault(err) {
		status = http.StatusInternalServerError
	}
	SendError(ctx, status, err)
}
