package main

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"rest-api/docs"
	"rest-api/logging"
	"rest-api/models"
	"rest-api/routes"
)

func setupRouter() *gin.Engine {
	router := gin.New()         // basic gin router
	router.Use(gin.Recovery())  // recover form panics and answer with 500
	router.Use(requestid.New()) // generate a unique id for every request (for logging)

	router.Use(logger.SetLogger(logger.WithLogger(logging.GinLogger)))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/v1")

	// this will expose GET /v1/healthz
	v1.GET("/healthz", routes.HealthzRoute)

	v1.PUT("/er", routes.CreateEmergencyRoom)
	v1.GET("/er", routes.GetEmergencyRooms)
	v1.GET("/er/{id}", routes.GetEmergencyRoomById)
	v1.PATCH("/er/{id}", routes.UpdateEmergencyRoom)
	v1.DELETE("/er/{id}", routes.DeleteEmergencyRoom)

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

	logging.SetupLogging(
		GetEnvOr("GIN_MODE", "development"),
		GetEnvOr("LOG_LEVEL", "info"),
	)

	if dotenvErr != nil {
		log.Fatal().Err(dotenvErr).Msg("Error loading .env file: ")
	} else {
		log.Info().Msg("no error loading .env file")
	}

	models.SetupDatabase(
		GetEnvOr("POSTGRES_HOST", "localhost"),
		GetEnvOr("POSTGRES_USER", "postgres"),
		GetEnvOr("POSTGRES_PASSWORD", "postgres"),
		GetEnvOr("POSTGRES_DB", "postgres"),
		GetEnvOr("POSTGRES_PORT", "5432"),
	)

	gin.SetMode(GetEnvOr("GIN_MODE", "debug"))

	setSwaggerInfo()

	router := setupRouter()

	addr := ":" + GetEnvOr("PORT", "3000")
	log.Info().Str("addr", addr).Msg("starting server")
	if err := router.Run(addr); err != nil {
		log.Fatal().Err(err).Msg("Could not start server.")
	}
}
