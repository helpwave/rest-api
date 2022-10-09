package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"rest-api/docs"
	"rest-api/models"
	"rest-api/routes"
)

func setupLogging() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if GetEnvOr("GIN_MODE", "development") != "release" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
	log.Logger = log.With().Caller().Logger()
	level, err := zerolog.ParseLevel(GetEnvOr("LOG_LEVEL", "info"))
	if err != nil {
		log.Fatal().Err(err).Msg("could not parse LOG_LEVEL")
	}
	log.Logger = log.Level(level)
	log.Info().Msg("Logging is set up")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

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
	setupLogging()

	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file: ")
	}
	log.Info().Msg("no error loading .env file")

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

	if err := router.Run(":" + GetEnvOr("PORT", "3000")); err != nil {
		log.Fatal().Err(err).Msg("Could not start server.")
	}
}
