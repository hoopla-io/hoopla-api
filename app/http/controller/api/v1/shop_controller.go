package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	shop_request "github.com/qahvazor/qahvazor/app/http/request/shop"
	"github.com/qahvazor/qahvazor/app/http/response"
	"github.com/qahvazor/qahvazor/internal/service"
)

type ShopController struct {
	service service.ShopService
}

func NewShopController(service service.ShopService) *ShopController {
	return &ShopController{
		service: service,
	}
}

// @Tags Shop
// @Accept  json
// @Produce  json
// @Param data body shop_request.GetShopDetailsRequest true "Get Shop Details Request"
// @Success 200 {array} shop_response.GetShopDetailsResponse "Shop details response"
// @Router /shop/detail [post]
func (ctr *ShopController) GetShopDetails(ctx *gin.Context) {
	getShopDetailsRequest := shop_request.GetShopDetailsRequest{}
	if err := ctx.ShouldBind(&getShopDetailsRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.GetShopDetails(getShopDetailsRequest)
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}