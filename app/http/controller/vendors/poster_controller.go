package vendor_controllers

import (
	"fmt"
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

// @Tags Poster
// @Accept  json
// @Produce  json
// @Param data query vendors_poster_request.OauthRequest true "OAUTH for poster"
// @Router /vendors/poster/oauth [get]
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

// @Tags Poster
// @Accept  json
// @Produce  json
// @Param data body vendors_poster_request.WebhookRequest true "Webhook for poster"
// @Router /vendors/poster/webhook [post]
func (c *PosterController) Webhook(ctx *gin.Context) {
	request := vendors_poster_request.WebhookRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	//verify := fmt.Sprintf(
	//	"%s;%s;%s;%s",
	//	request.Account,
	//	request.Object,
	//	request.ObjectID,
	//	request.Action,
	//)
	//verify = fmt.Sprintf("%x", md5.Sum([]byte(verify)))

	//object->incoming_order
	fmt.Println(request)
}
