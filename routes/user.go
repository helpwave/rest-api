package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"rest-api/models"
)

// CreateUserRoute godoc
// @summary Create a user
// @Router /user [put]
func CreateUserRoute(ctx *gin.Context) {
	if models.DB == nil {
		log.Fatalln("db not set")
	}
	res := models.DB.Create(&models.User{})
	if res.Error != nil {
		log.Println(res.Error)
		ctx.JSON(500, "Database issues :/")
	} else {
		ctx.JSON(200, gin.H{})
	}
}

// GetUsersRoute godoc
// @summary Get all user
// @Router /user [get]
func GetUsersRoute(ctx *gin.Context) {
	if models.DB == nil {
		log.Fatalln("db not set")
	}
	var users []models.User
	res := models.DB.Find(&users)
	if res.Error != nil {
		log.Println(res.Error)
		ctx.JSON(500, "Database issues :/")
	}
	ctx.JSON(200, users)
}
