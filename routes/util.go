package routes

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"rest-api/models"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type HTTPError struct {
	Code    int    `example:"500"`
	Message string `example:"Some complicated error message here"`
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

func GetParamUUID(ctx *gin.Context, param string) (uuid.UUID, error) {
	raw := ctx.Param(param)
	parsed, err := uuid.Parse(raw)
	if err != nil {
		return uuid.UUID{}, errors.New(raw + " is an invalid uuid")
	}
	return parsed, nil
}
