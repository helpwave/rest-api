package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"rest-api/logging"
	"rest-api/models"
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

// DeleteDepartment godoc
// @Summary    delete a department
// @Tags       departments
// @Produce    json
// @Param      id                              path        uuid                    true    "department id"
// @Success    200							   {object}    HTTPSuccessResponse
// @Failure    400                             {object}    HTTPErrorResponse
// @Failure    404                             {object}    HTTPErrorResponse
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

	db.Delete(&models.Department{
		DepartmentBase: models.DepartmentBase{
			ID: departmentId,
		},
	})

	if err != nil {
		HandleDBError(ctx, logCtx, err)
		return
	}

	ctx.Status(http.StatusOK)
}
