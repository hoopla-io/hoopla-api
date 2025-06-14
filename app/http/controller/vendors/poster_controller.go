package vendor_controllers

import (
	"github.com/gin-gonic/gin"
	vendors_poster_request "github.com/hoopla/hoopla-api/app/http/request/vendors/poster"
	"github.com/hoopla/hoopla-api/app/http/response"
	"github.com/hoopla/hoopla-api/internal/service"
	vendor_utils "github.com/hoopla/hoopla-api/utils/vendors"
	"net/http"
)

type PosterController struct {
	partnerService      service.PartnerService
	partnerTokenService service.PartnerTokenService
}

func NewPosterController(partnerService service.PartnerService, partnerTokenService service.PartnerTokenService) *PosterController {
	return &PosterController{
		partnerService:      partnerService,
		partnerTokenService: partnerTokenService,
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

	if request.Object == "incoming_order" {
		partner, code, err := c.partnerService.GetPartnerByVendorId(request.Account)
		if err != nil {
			if code == http.StatusNotFound {
				response.NotFoundResponse(ctx, request.Account)
				return
			}
			response.ErrorResponse(ctx, code, err.Error())
			return
		}

		if request.Action == "changed" {
			poster := vendor_utils.Poster{
				VendorID:    partner.VendorID,
				VendorKey:   partner.VendorKey,
				AccessToken: "",
			}

			accessToken, err := c.partnerTokenService.GetAccessToken(partner)
			if err != nil {
				response.ErrorResponse(ctx, 500, err.Error())
				return
			}
			poster.AccessToken = accessToken

			poster.GetOrderStatus(request.ObjectID)
		}
	}

	//object->incoming_order
	//fmt.Println(request)
}
