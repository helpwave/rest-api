package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/models"
)

type GetSingleERResponse struct {
	models.EmergencyRoom
	Departments []models.DepartmentBase
}

// GetEmergencyRoomById godoc
// @Summary		get an emergency room by id
// @Tags 		emergency-rooms
// @Produce 	json
// @Param       id   		path 		string				true	"Emergency Room's ID"
// @Success     200  		{object} 	GetSingleERResponse
// @Failure     501  		{object}	HTTPErrorResponse
// @Router      /er/{id} 	[get]
func GetEmergencyRoomById(ctx *gin.Context) {
	_ = ctx.Param("id")
	SendError(ctx, http.StatusNotImplemented, errors.New("this endpoint is not implemented yet"))
}

type PutERRequest struct {
	models.EmergencyRoomBase
	Departments []models.DepartmentBase
}

// CreateEmergencyRoom godoc
// @Summary     create a new emergency room
// @Tags 		emergency-rooms
// @Accept	 	json
// @Produce 	json
// @Param		emergency-room 	body 		PutERRequest		true	"ER to add"
// @Success     200  			{object} 	GetSingleERResponse
// @Failure     501  			{object}	HTTPErrorResponse
// @Router      /er/{id} 		[put]
func CreateEmergencyRoom(ctx *gin.Context) {
	SendError(ctx, http.StatusNotImplemented, errors.New("this endpoint is not implemented yet"))
}

// UpdateEmergencyRoom godoc
// @Summary     update an emergency room by id
// @Tags 		emergency-rooms
// @Produce 	json
// @Param       id   		path 		string				true	"Emergency Room's ID"
// @Success     200  		{object} 	GetSingleERResponse
// @Failure     501  		{object}	HTTPErrorResponse
// @Router      /er/{id}	[patch]
func UpdateEmergencyRoom(ctx *gin.Context) {
	_ = ctx.Param("id")
	SendError(ctx, http.StatusNotImplemented, errors.New("this endpoint is not implemented yet"))
}

// DeleteEmergencyRoom godoc
// @Summary      delete an emergency room by id
// @Tags 		 emergency-rooms
// @Produce 	 json
// @Param        id   path 		string  			true	"Emergency Room's ID"
// @Success      200  {object} 	StatusResponse
// @Failure      501  {object}  HTTPErrorResponse
// @Router       /er/{id} [delete]
func DeleteEmergencyRoom(ctx *gin.Context) {
	_ = ctx.Param("id")
	SendError(ctx, http.StatusNotImplemented, errors.New("this endpoint is not implemented yet"))
}

type GetMultipleERsResponse struct {
	PaginatedResponse
	EmergencyRooms []GetSingleERResponse
}

// GetEmergencyRooms godoc
// @Summary      get emergency rooms
// @Tags 		 emergency-rooms
// @Produce 	 json
// @Param        page	query		uint					false	"0-indexed page number, 0 is assumed when omitted"
// @Success      200 	{object}	GetMultipleERsResponse
// @Failure      501 	{object}	HTTPErrorResponse
// @Router       /er	[get]
func GetEmergencyRooms(ctx *gin.Context) {
	_ = ctx.Param("page")
	SendError(ctx, http.StatusNotImplemented, errors.New("this endpoint is not implemented yet"))
}