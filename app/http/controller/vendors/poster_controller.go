package vendor_controllers

import (
	"github.com/gin-gonic/gin"
	vendors_poster_request "github.com/hoopla/hoopla-api/app/http/request/vendors/poster"
	"github.com/hoopla/hoopla-api/app/http/response"
	"github.com/hoopla/hoopla-api/internal/service"
)

type PosterController struct {
	partnerService service.PartnerService
}

func NewPosterController(partnerService service.PartnerService) *PosterController {
	return &PosterController{
		partnerService: partnerService,
	}
}

func (c *PosterController) Oauth(ctx *gin.Context) {
	request := vendors_poster_request.OauthRequest{}
	if err := ctx.ShouldBindQuery(&request); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	code, err := c.partnerService.UpdateVendorKey(request.Account, request.Code)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, "OK!", nil, nil)
}

func (c *PosterController) Webhook(ctx *gin.Context) {}
