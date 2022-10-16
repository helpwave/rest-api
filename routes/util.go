package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

// StatusResponse should give the caller feedback on the state of the task they submitted
type StatusResponse struct {
	Ok bool
}

// SendOk sends a basic "OK" response
func SendOk(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, StatusResponse{Ok: true})
}
