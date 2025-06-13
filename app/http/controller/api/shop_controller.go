package api

import (
	"github.com/gin-gonic/gin"
	shops_request "github.com/hoopla/hoopla-api/app/http/request/shops"
	"github.com/hoopla/hoopla-api/app/http/response"
	"github.com/hoopla/hoopla-api/internal/service"
)

type ShopController struct {
	shopService service.ShopService
}

func NewShopController(shopService service.ShopService) *ShopController {
	return &ShopController{
		shopService: shopService,
	}
}

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

// @Tags Shops
// @Accept  json
// @Produce  json
// @Param data query shops_request.NearShopsRequest true "Near shops"
// @Router /v1/shops/near-shops [get]
func (c *ShopController) NearShops(ctx *gin.Context) {
	var request shops_request.NearShopsRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	shops, code, err := c.shopService.NearShopsList(request)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, "ok!", shops, nil)
	return
}

// @Tags Shops
// @Accept  json
// @Produce  json
// @Param data query shops_request.ShopRequest true "Partner shops"
// @Router /v1/shops/shop [get]
func (c *ShopController) Shop(ctx *gin.Context) {
	var request shops_request.ShopRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	shop, code, err := c.shopService.ShopDetail(request)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, "ok!", shop, nil)
	return
}
