package api

import (
	"github.com/gin-gonic/gin"
	partners_request "github.com/hoopla/hoopla-api/app/http/request/partners"
	"github.com/hoopla/hoopla-api/app/http/response"
	"github.com/hoopla/hoopla-api/internal/service"
)

type PartnerController struct {
	partnerService service.PartnerService
}

func NewPartnerController(PartnerService service.PartnerService) *PartnerController {
	return &PartnerController{
		partnerService: PartnerService,
	}
}

func (c PartnerController) Partners(ctx *gin.Context) {
	var request partners_request.PartnersRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
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

func (c PartnerController) Partner(ctx *gin.Context) {
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
