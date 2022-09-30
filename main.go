package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")

	// this will expose GET /v1/healthz
	v1.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"server": "ok",
		})
	})
	return router
}

func main() {
	dotenvErr := godotenv.Load()
	if dotenvErr != nil {
		log.Fatalln("Error loading .env file: ", dotenvErr)
	}

	router := setupRouter()

	serverErr := router.Run(":" + GetEnvOr("PORT", "3000"))
	if serverErr != nil {
		log.Fatalln("Could not start server; See logs why.")
	}
}
