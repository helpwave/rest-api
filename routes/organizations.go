package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/logging"
	"rest-api/models"
	"rest-api/util"
)

type CreateOrgRequest struct {
	LongName     string `binding:"required" example:"Uniklinikum Münster"`
	ShortName    string `gorm:"default:NULL" example:"UKM"`
	ContactEmail string `binding:"required,email" example:"example@helpwave.de"`
}

type GetSingleOrgResponse struct {
	models.OrganizationBase
}

// CreateOrganization godoc
// @Summary    create a new organization
// @Tags       organizations
// @Accept     json
// @Produce    json
// @Param      authorization                   header      string                     true    "Bearer: <TOKEN>"
// @Param      organization                    body        CreateOrgRequest           true    "Org to add"
// @Success    200                             {object}    GetSingleOrgResponse
// @Failure    400                             {object}    HTTPErrorResponse
// @Router     /organizations                  [put]
func CreateOrganization(ctx *gin.Context) {
	log, logCtx := logging.GetRequestLogger(ctx)

	//
	// Validate body
	//
	body := CreateOrgRequest{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Warn().Err(err).Msg("validation failed")
		SendError(ctx, http.StatusBadRequest, err)
		return
	}
	log.Debug().Str("body", util.Formatted(body)).Send()

	//
	// create model for gORM
	//
	orga := models.Organization{
		OrganizationBase: models.OrganizationBase{
			LongName:     body.LongName,
			ShortName:    body.ShortName,
			ContactEmail: body.ContactEmail,
		},
	}
	log.Debug().Str("model", util.Formatted(orga)).Send()

	db := models.GetDB(logCtx)

	res := db.Create(&orga)
	if err := res.Error; err != nil {
		HandleDBError(ctx, logCtx, err)
		return
	}

	resp := GetSingleOrgResponse{
		OrganizationBase: orga.OrganizationBase,
	}

	ctx.JSON(http.StatusOK, resp)
}
