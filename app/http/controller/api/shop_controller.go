package api

import (
	"github.com/gin-gonic/gin"
	shops_request "github.com/qahvazor/qahvazor/app/http/request/shops"
	"github.com/qahvazor/qahvazor/app/http/response"
	"github.com/qahvazor/qahvazor/internal/service"
)

type ShopController struct {
	shopService service.ShopService
}

func NewShopController(shopService service.ShopService) *ShopController {
	return &ShopController{
		shopService: shopService,
	}
}

// @Tags Partners
// @Accept  json
// @Produce  json
// @Param data query shops_request.PartnerShopsRequest true "Partner shops"
// @Router /shops/partner-shops [get]
func (c *ShopController) PartnerShopList(ctx *gin.Context) {
	var request shops_request.PartnerShopsRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	shops, code, err := c.shopService.PartnerShopsList(request)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, "ok!", shops, nil)
	return
}
