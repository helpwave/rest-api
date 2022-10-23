package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// VersionRoute godoc
// @Summary      version route
// @Description  return version of the running
// @Success      200                            {string}  string
// @Router       /version                       [get]
func VersionRoute(version string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, version)
	}
}
