package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/logging"
	"rest-api/models"
	"rest-api/util"
)

type GetDepartmentsResponse struct {
	PaginatedResponse
	Departments []models.DepartmentBase
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

type SingleDepartmentResponse struct {
	models.DepartmentBase
}

type CreateDepartmentRequest struct {
	Name string `binding:"required"`
}

// CreateDepartment godoc
// @Summary    create new department
// @Tags       departments
// @Produce    json
// @Param      authorization                   header      string                    true    "Bearer: <TOKEN>"
// @Param      department                      body        CreateDepartmentRequest   true    "Dep. to add"
// @Success    200                             {object}    SingleDepartmentResponse
// @Failure    400                             {object}    HTTPErrorResponse
// @Router     /department                     [put]
func CreateDepartment(ctx *gin.Context) {
	log, logCtx := logging.GetRequestLogger(ctx)
	db := models.GetDB(logCtx)

	//
	// Validate body
	//
	body := CreateDepartmentRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Warn().Err(err).Msg("validation failed")
		SendError(ctx, http.StatusBadRequest, err)
		return
	}
	log.Debug().Str("body", util.Formatted(body)).Send()

	//
	// create model for gORM
	//
	department := models.Department{
		DepartmentBase: models.DepartmentBase{Name: body.Name},
	}
	log.Debug().Str("model", util.Formatted(department)).Send()

	//
	// add to database
	//
	res := db.Create(&department)
	if err := res.Error; err != nil {
		HandleDBError(ctx, logCtx, err)
		return
	}

	//
	// return result
	//
	resp := SingleDepartmentResponse{
		DepartmentBase: department.DepartmentBase,
	}

	ctx.JSON(http.StatusOK, resp)
}
