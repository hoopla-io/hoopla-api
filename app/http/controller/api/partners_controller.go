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

func NewPartnersController(PartnerService service.PartnerService) *PartnersController {
	return &PartnersController{
		partnerService: PartnerService,
	}
}

// @Tags Partners
// @Accept  json
// @Produce  json
// @Param data query partners_request.PartnersRequest true "Partners list"
// @Router /partners [get]
func (c PartnersController) Partners(ctx *gin.Context) {
	var request partners_request.PartnersRequest
	if err := ctx.ShouldBind(&request); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	partners, code, err := c.partnerService.PartnersList(request)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, "ok!", partners, nil)
	return
}

// @Tags Partners
// @Accept  json
// @Produce  json
// @Param data query partners_request.PartnerRequest true "Partner Detail"
// @Router /partners/partner [get]
func (c PartnersController) Partner(ctx *gin.Context) {
	var request partners_request.PartnerRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	partner, code, err := c.partnerService.PartnerDetail(request)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, "ok!", partner, nil)
	return
}
