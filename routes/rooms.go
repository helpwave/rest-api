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
	ID uuid.UUID
	models.EmergencyRoomBase
	Departments []models.DepartmentBase
}

// GetEmergencyRoomById godoc
// @Summary get an emergency room by id
// @Tags    emergency-rooms
// @Produce json
// @Param   id              path        string                true    "Emergency Room's ID"
// @Success 200             {object}    GetSingleERResponse
// @Failure 400             {object}    HTTPErrorResponse
// @Router  /er/{id}        [get]
func GetEmergencyRoomById(ctx *gin.Context) {
	log, logCtx := logging.GetRequestLogger(ctx)
	db := models.GetDB(logCtx)

	erIdRaw := ctx.Param("id")
	log.Debug().Str("requested_id", erIdRaw).Send()
	erID, err := uuid.Parse(erIdRaw)
	if err != nil {
		SendError(ctx, http.StatusBadRequest, errors.New("invalid uuid"))
		return
	}

	er := models.EmergencyRoom{
		ID: erID,
	}

	tx := db.First(&er)
	if tx.Error != nil {
		HandleDBError(ctx, logCtx, tx.Error)
		return
	}

	resp := GetSingleERResponse{
		ID:                er.ID,
		EmergencyRoomBase: er.EmergencyRoomBase,
		Departments:       models.DepartmentsToBases(er.Departments),
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetEmergencyRooms godoc
// @Summary    get all emergency rooms
// @Tags       emergency-rooms
// @Produce    json
// @Param      page                            query       uint                    false   "0-indexed page number, 0 is assumed when omitted"
// @Param      page_size                       query       uint					   false   "page size, 100 is assumed when omitted"
// @Success    200                             {object}    GetMultipleERsResponse
// @Failure    400                             {object}    HTTPErrorResponse
// @Router     /er                             [get]
func GetEmergencyRooms(ctx *gin.Context) {
	_, logCtx := logging.GetRequestLogger(ctx)
	db := models.GetDB(logCtx)

	pagination, err := GetPagination(ctx, db, models.EmergencyRoom{})
	if err != nil {
		HandleDBError(ctx, logCtx, err)
		return
	}

	var emergencyRooms []models.EmergencyRoom
	tx := db.Scopes(Paginate(pagination)).Select("id").Find(&emergencyRooms)

	if err := tx.Error; err != nil {
		HandleDBError(ctx, logCtx, err)
		return
	}

	ids := make([]uuid.UUID, len(emergencyRooms))
	for i, emergencyRoom := range emergencyRooms {
		ids[i] = emergencyRoom.ID
	}

	resp := GetMultipleERsResponse{
		PaginatedResponse: pagination,
		EmergencyRooms:    ids,
	}

	ctx.JSON(http.StatusOK, resp)
}

type PutERRequest struct {
	models.EmergencyRoomBase
	Departments []uuid.UUID
}

// CreateEmergencyRoom godoc
// @Summary    create a new emergency room
// @Tags       emergency-rooms
// @Accept     json
// @Produce    json
// @Param      authorization                   header      string                true    "Bearer: <TOKEN>"
// @Param      emergency-room                  body        PutERRequest          true    "ER to add"
// @Success    200                             {object}    GetSingleERResponse
// @Failure    400                             {object}    HTTPErrorResponse
// @Failure    501                             {object}    HTTPErrorResponse
// @Router     /er                             [put]
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
	log.Debug().Str("body", util.Formatted(body)).Send()

	//
	// create model for gORM
	//
	er := models.EmergencyRoom{
		EmergencyRoomBase: body.EmergencyRoomBase,
		Departments:       models.UUIDsToDepartments(body.Departments),
	}
	log.Debug().Str("model", util.Formatted(er)).Send()

	db := models.GetDB(logCtx)
	db = db.Omit("Departments.*") // do not attempt to create ("upsert") Departments, they have to exist already

	res := db.Create(&er)
	if err := res.Error; err != nil {
		HandleDBError(ctx, logCtx, err)
		return
	}

	resp := GetSingleERResponse{
		ID:                er.ID,
		EmergencyRoomBase: er.EmergencyRoomBase,
		Departments:       models.DepartmentsToBases(er.Departments),
	}

	ctx.JSON(http.StatusOK, resp)
}

// UpdateEmergencyRoom godoc
// @Summary    update an emergency room by id
// @Tags       emergency-rooms
// @Produce    json
// @Param      authorization                   header      string                true    "Bearer: <TOKEN>"
// @Param      id                              path        string                true    "Emergency Room's ID"
// @Param      emergency-room                  body        PutERRequest          true    "ER to update"
// @Success    200
// @Failure    400                             {object}    HTTPErrorResponse
// @Router     /er/{id}                        [patch]
func UpdateEmergencyRoom(ctx *gin.Context) {
	log, logCtx := logging.GetRequestLogger(ctx)
	db := models.GetDB(logCtx)

	erIdRaw := ctx.Param("id")
	log.Debug().Str("update_id", erIdRaw).Send()
	erID, err := uuid.Parse(erIdRaw)
	if err != nil {
		SendError(ctx, http.StatusBadRequest, errors.New("invalid uuid"))
		return
	}

	er := models.EmergencyRoom{
		ID: erID,
	}

	// validate update body
	body := PutERRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Warn().Err(err).Msg("validation failed")
		SendError(ctx, http.StatusBadRequest, err)
		return
	}
	log.Debug().Str("body", util.Formatted(body)).Send()

	// create the updating model for gORM
	updatedEr := models.EmergencyRoom{
		EmergencyRoomBase: body.EmergencyRoomBase,
		Departments:       models.UUIDsToDepartments(body.Departments),
	}
	log.Debug().Str("model", util.Formatted(er)).Send()

	// this performs the actual update
	tx := db.Where(&er).Updates(&updatedEr)
	if tx.Error != nil {
		HandleDBError(ctx, logCtx, tx.Error)
		return
	}

	ctx.Status(http.StatusOK)
}

// DeleteEmergencyRoom godoc
// @Summary    delete an emergency room by id
// @Tags       emergency-rooms
// @Produce    json
// @Param      authorization                   header      string                true    "Bearer: <TOKEN>"
// @Param      id                              path        string                true    "Emergency Room's ID"
// @Success    200
// @Failure    400                             {object}    HTTPErrorResponse
// @Router    /er/{id}                         [delete]
func DeleteEmergencyRoom(ctx *gin.Context) {
	log, logCtx := logging.GetRequestLogger(ctx)
	db := models.GetDB(logCtx)

	erIdRaw := ctx.Param("id")
	log.Debug().Str("delete_id", erIdRaw).Send()
	erID, err := uuid.Parse(erIdRaw)
	if err != nil {
		SendError(ctx, http.StatusBadRequest, errors.New("invalid uuid"))
		return
	}

	er := models.EmergencyRoom{
		ID: erID,
	}

	tx := db.Delete(&er)
	if tx.Error != nil {
		HandleDBError(ctx, logCtx, tx.Error)
		return
	}

	ctx.Status(http.StatusOK)
}

type GetMultipleERsResponse struct {
	PaginatedResponse
	EmergencyRooms []uuid.UUID
}
