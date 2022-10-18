package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthzResponse items are either "ok" or "err"
type HealthzResponse struct {
	Server string `json:"server"`
}

// HealthzRoute godoc
// @Summary      health check route
// @Description  can be used for health checks
// @Success      200  {object}     HealthzResponse
// @Failure      500  {object}     HealthzResponse
// @Router       /healthz [get]
func HealthzRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, HealthzResponse{
		Server: "ok",
	})
}
