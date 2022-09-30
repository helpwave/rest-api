package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"rest-api/docs"
	"rest-api/routes"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/v1")

	// this will expose GET /v1/healthz
	v1.GET("/healthz", routes.HealthzRoute)
	return router
}

func setSwaggerInfo() {
	docs.SwaggerInfo.Title = "helpwave rest-api"
	docs.SwaggerInfo.Description = "helpwave rest-api backend"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "main.helpwave.de"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"https"}
}

func main() {
	dotenvErr := godotenv.Load()
	if dotenvErr != nil {
		log.Fatalln("Error loading .env file: ", dotenvErr)
	}

	gin.SetMode(GetEnvOr("GIN_MODE", "debug"))

	setSwaggerInfo()

	router := setupRouter()

	serverErr := router.Run(":" + GetEnvOr("PORT", "3000"))
	if serverErr != nil {
		log.Fatalln("Could not start server; See logs why.")
	}
}
