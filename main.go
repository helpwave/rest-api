package main

import (
	"rest-api/docs"
	"rest-api/logging"
	"rest-api/models"
	"rest-api/routes"
	"rest-api/util"

	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var Version string

func setupRouter() *gin.Engine {
	router := gin.New()        // basic gin router
	router.Use(gin.Recovery()) // recover form panics and answer with 500

	router.Use(requestid.New())                                        // generate a unique id for every request (for logging)
	router.Use(logger.SetLogger(logger.WithLogger(logging.GinLogger))) // set gin's request logger

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/v1")

	// this will expose GET /v1/healthz
	v1.GET("/healthz", routes.HealthzRoute)
	v1.GET("/version", routes.VersionRoute(Version))

	v1.PUT("/er", routes.AuthMiddleware(), routes.CreateEmergencyRoom)
	v1.GET("/er", routes.GetEmergencyRooms)
	v1.GET("/er/:id", routes.GetEmergencyRoomById)
	v1.PATCH("/er/:id", routes.AuthMiddleware(), routes.UpdateEmergencyRoom)
	v1.DELETE("/er/:id", routes.AuthMiddleware(), routes.DeleteEmergencyRoom)

	v1.GET("/departments", routes.GetDepartments)
	v1.PATCH("/department/:id", routes.AuthMiddleware(), routes.UpdateDepartment)
	return router
}

func setSwaggerInfo() {
	docs.SwaggerInfo.Title = "helpwave rest-api"
	docs.SwaggerInfo.Description = "helpwave rest-api backend"
	docs.SwaggerInfo.Version = Version
	docs.SwaggerInfo.Host = util.GetEnvOr("BASE_URI", "api.helpwave.de")
	docs.SwaggerInfo.BasePath = "/v1"

	scheme := "https"
	if util.GetEnvOr("SCHEME_HTTPS", "true") == "false" {
		scheme = "http"
	}
	docs.SwaggerInfo.Schemes = []string{scheme}
}

func main() {
	dotenvErr := godotenv.Load()

	GinMode := util.GetEnvOr("GIN_MODE", "development")
	LogLevel := util.GetEnvOr("LOG_LEVEL", "info")

	logging.SetupLogging(
		GinMode,
		LogLevel,
		Version,
	)

	if len(Version) == 0 && GinMode != "development" {
		log.Warn().Msg("Version is empty in production build! Recompile using ldflag '-X main.Version=<version>'")
	}

	if dotenvErr == nil {
		log.Info().Msg("successfully loaded .env file")
	}

	models.SetupDatabase(
		util.GetEnvOr("POSTGRES_HOST", "localhost"),
		util.GetEnvOr("POSTGRES_USER", "postgres"),
		util.GetEnvOr("POSTGRES_PASSWORD", "postgres"),
		util.GetEnvOr("POSTGRES_DB", "postgres"),
		util.GetEnvOr("POSTGRES_PORT", "5432"),
	)

	gin.SetMode(util.GetEnvOr("GIN_MODE", "debug"))

	setSwaggerInfo()

	router := setupRouter()

	addr := ":" + util.GetEnvOr("PORT", "3000")
	log.Info().Str("addr", addr).Msg("starting server")
	if err := router.Run(addr); err != nil {
		log.Fatal().Err(err).Msg("Could not start server.")
	}
}
