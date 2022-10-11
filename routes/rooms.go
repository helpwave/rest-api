package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rest-api/logging"
	"rest-api/models"
	"rest-api/util"
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
	Departments []uuid.UUID
}

// CreateEmergencyRoom godoc
// @Summary     create a new emergency room
// @Tags 		emergency-rooms
// @Accept	 	json
// @Produce 	json
// @Param		emergency-room 	body 		PutERRequest		true	"ER to add"
// @Success     200  			{object} 	GetSingleERResponse
// @Failure		400				{object}	HTTPErrorResponse
// @Failure     501  			{object}	HTTPErrorResponse
// @Router      /er				[put]
func CreateEmergencyRoom(ctx *gin.Context) {
	log, logCtx := logging.GetRequestLogger(ctx)

	//
	// Validate body
	//
	body := PutERRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Warn().Err(err).Msg("validation failed")
		SendError(ctx, http.StatusBadRequest, err)
		return
	}
	log.Debug().Msgf("req body:", util.Formatted(body))

	//
	// convert department UUIDs into Departments
	//
	deps := make([]models.Department, len(body.Departments))
	for i := range body.Departments {
		deps[i].ID = body.Departments[i]
	}

	//
	// create model for gORM
	//
	er := models.EmergencyRoom{
		EmergencyRoomBase: body.EmergencyRoomBase,
		Departments:       deps,
	}
	log.Debug().Msgf("model", util.Formatted(er))

	db := models.GetDB(logCtx)
	db = db.Omit("Departments.*") // do not attempt to create ("upsert") Departments, they have to exist already

	res := db.Create(&er)
	if err := res.Error; err != nil {
		HandleDBError(ctx, logCtx, err)
		return
	}

	resp := GetSingleERResponse{
		EmergencyRoom: er,
		Departments:   models.DepartmentsToBases(er.Departments),
	}

	ctx.JSON(http.StatusOK, resp)
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
