package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"rest-api/logging"
	"rest-api/models"
)

type CreateUserRequest struct {
	Email        string `binding:"required,email" example:"example@helpwave.de"`
	FullName     string `binding:"required" example:"Some Name"`
	Password     string `binding:"required,min=6,max=100" example:"hunter2"`
	Admin        bool
	Organization uuid.UUID
}

type CreateUserResponse struct {
	UserID uuid.UUID `json:"userID"`
}

// CreateUser godoc
// @Summary    create a new user
// @Tags       users
// @Accept     json
// @Produce    json
// @Param      authorization    header      string                 true    "Bearer: <TOKEN>"
// @Param      user             body        CreateUserRequest      true    "user to add"
// @Success    200              {object}    CreateUserResponse
// @Failure    400              {object}    HTTPErrorResponse
// @Failure    500              {object}    HTTPErrorResponse
// @Router     /users           [put]
func CreateUser(ctx *gin.Context) {
	log, logCtx := logging.GetRequestLogger(ctx)

	//
	// Validate body
	//
	body := CreateUserRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Warn().Err(err).Msg("validation failed")
		SendError(ctx, http.StatusBadRequest, err)
		return
	}

	if body.Organization == uuid.Nil && !body.Admin {
		log.Warn().Msg("Attempt to create non-admin user without organization")
		SendError(ctx, http.StatusBadRequest, errors.New("no organization specified for non-admin user"))
		return
	}

	// TODO: additional password complexity enforcement

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("could not calculate bcrypt hash")
		SendError(ctx, http.StatusInternalServerError, errors.New("an error occurred while generating the password hash"))
		return
	}

	//
	// create model for gORM
	//
	user := models.User{
		UserBase: models.UserBase{
			Email:          body.Email,
			FullName:       body.FullName,
			OrganizationID: body.Organization,
		},
		PwBcrypt: string(hashBytes),
	}

	db := models.GetDB(logCtx)
	res := db.Create(&user)

	if err := res.Error; err != nil {
		HandleDBError(ctx, logCtx, err)
		return
	}

	resp := CreateUserResponse{
		UserID: user.ID,
	}

	ctx.JSON(http.StatusOK, resp)
}
