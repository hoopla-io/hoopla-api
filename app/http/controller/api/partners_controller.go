package api

import (
	"github.com/gin-gonic/gin"
	partners_request "github.com/qahvazor/qahvazor/app/http/request/partners"
	"github.com/qahvazor/qahvazor/app/http/response"
	"github.com/qahvazor/qahvazor/internal/service"
)

type PartnersController struct {
	partnerService service.PartnerService
}

func NewPartnersController(partnerService service.PartnerService) *PartnersController {
	return &PartnersController{
		partnerService: partnerService,
	}
}

// @Tags Partners
// @Accept  json
// @Produce  json
// @Param data body partners_request.PartnersRequest true "Partners list"
// @Router /partners [get]
func (c PartnersController) Partners(ctx *gin.Context) {
	var request partners_request.PartnersRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

}

// @Tags Partners
// @Accept  json
// @Produce  json
// @Param data body partners_request.PartnerRequest true "Partner Detail"
// @Router /partners/partner [get]
func (c PartnersController) Partner(ctx *gin.Context) {
	var request partners_request.PartnerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}
}
