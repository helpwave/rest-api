package routes

import (
	"net/http"
	"rest-api/logging"
	"rest-api/models"
	"rest-api/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetDepartmentsResponse struct {
	PaginatedResponse
	Departments []models.DepartmentBase `json:"departments"`
}

// GetDepartments godoc
// @Summary    get all departments
// @Tags       departments
// @Produce    json
// @Param      page                            query       uint                    false   "0-indexed page number, 0 is assumed when omitted"
// @Param      page_size                       query       uint                    false   "page size, 100 is assumed when omitted"
// @Success    200                             {object}    GetDepartmentsResponse
// @Failure    400                             {object}    HTTPErrorResponse
// @Router     /departments                    [get]
func GetDepartments(ctx *gin.Context) {
	_, logCtx := logging.GetRequestLogger(ctx)
	db := models.GetDB(logCtx)

	pagination, err := GetPagination(ctx, db, models.Department{})
	if err != nil {
		HandleDBError(ctx, logCtx, err)
		return
	}

	var departments []models.Department
	tx := db.Scopes(Paginate(pagination)).Select("id, name").Find(&departments)

	if err := tx.Error; err != nil {
		HandleDBError(ctx, logCtx, err)
		return
	}

	resp := GetDepartmentsResponse{
		PaginatedResponse: pagination,
		Departments:       models.DepartmentsToBases(departments),
	}

	ctx.JSON(http.StatusOK, resp)
}

type UpdateDepartmentRequest struct {
	Name string `bind:"required" example:"pediatric surgery"`
}

// UpdateDepartment godoc
// @Summary    update an department by id
// @Tags       departments
// @Produce    json
// @Param      authorization                   header      string                   true    "Bearer: <TOKEN>"
// @Param      id                              path        string                   true    "Department's ID"
// @Param      department                      body        UpdateDepartmentRequest  true    "ER to update"
// @Success    200
// @Failure    400                             {object}    HTTPErrorResponse
// @Router     /departments/{id}                [patch]
func UpdateDepartment(ctx *gin.Context) {
	log, logCtx := logging.GetRequestLogger(ctx)
	db := models.GetDB(logCtx)

	id, err := GetParamUUID(ctx, "id")
	if err != nil {
		SendError(ctx, http.StatusBadRequest, err)
		return
	}

	department := models.Department{
		DepartmentBase: models.DepartmentBase{
			ID: id,
		},
	}
	log.Debug().Str("model", util.Formatted(department)).Send()

	// validate update body
	body := UpdateDepartmentRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Warn().Err(err).Msg("validation failed")
		SendError(ctx, http.StatusBadRequest, err)
		return
	}
	log.Debug().Str("body", util.Formatted(body)).Send()

	// create the updating model for gORM
	updatedDepartment := models.Department{
		DepartmentBase: models.DepartmentBase{
			Name: body.Name,
		},
	}

	// this performs the actual update
	tx := db.Where(&department).Updates(&updatedDepartment)
	if tx.Error != nil {
		HandleDBError(ctx, logCtx, tx.Error)
		return
	}

	ctx.Status(http.StatusOK)
}

type SingleDepartmentResponse struct {
	models.DepartmentBase
}

type CreateDepartmentRequest struct {
	Name string `binding:"required" example:"pediatric surgery"`
}

// CreateDepartment godoc
// @Summary    create new department
// @Tags       departments
// @Produce    json
// @Param      authorization                   header      string                    true    "Bearer: <TOKEN>"
// @Param      department                      body        CreateDepartmentRequest   true    "Dep. to add"
// @Success    200                             {object}    SingleDepartmentResponse
// @Failure    400                             {object}    HTTPErrorResponse
// @Router     /departments                    [put]
func CreateDepartment(ctx *gin.Context) {
	log, logCtx := logging.GetRequestLogger(ctx)
	db := models.GetDB(logCtx)

	// validate body
	body := CreateDepartmentRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Warn().Err(err).Msg("validation failed")
		SendError(ctx, http.StatusBadRequest, err)
		return
	}
	log.Debug().Str("body", util.Formatted(body)).Send()

	// create model for gORM
	department := models.Department{
		DepartmentBase: models.DepartmentBase{Name: body.Name},
	}
	log.Debug().Str("model", util.Formatted(department)).Send()

	// add to database
	res := db.Create(&department)
	if err := res.Error; err != nil {
		HandleDBError(ctx, logCtx, err)
		return
	}

	// return result
	resp := SingleDepartmentResponse{
		DepartmentBase: department.DepartmentBase,
	}

	ctx.JSON(http.StatusOK, resp)
}

// DeleteDepartment godoc
// @Summary    delete a department
// @Tags       departments
// @Produce    json
// @Param      authorization                   header      string                true    "Bearer: <TOKEN>"
// @Param      id                              path        string                true    "department id"
// @Success    200
// @Failure    400                             {object}    HTTPErrorResponse
// @Router     /departments/{id}               [delete]
func DeleteDepartment(ctx *gin.Context) {
	_, logCtx := logging.GetRequestLogger(ctx)
	db := models.GetDB(logCtx)

	departmentId, _ := uuid.Parse(ctx.Param("id"))

	err := db.Model(&models.Department{
		DepartmentBase: models.DepartmentBase{
			ID: departmentId,
		},
	}).Association("Rooms").Clear()

	tx := db.Delete(&models.Department{
		DepartmentBase: models.DepartmentBase{
			ID: departmentId,
		},
	})

	if tx.Error != nil {
		HandleDBError(ctx, logCtx, tx.Error)
		return
	}

	if err != nil {
		HandleDBError(ctx, logCtx, err)
		return
	}

	ctx.Status(http.StatusOK)
}
