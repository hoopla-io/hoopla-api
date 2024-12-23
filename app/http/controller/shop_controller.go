package controller

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

func (ctr *ShopController) CreateShop(ctx *gin.Context) {
	createShopRequest := shop_request.CreateShopRequest{}
	if err := ctx.ShouldBind(&createShopRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.CreateShop(createShopRequest)
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

func (ctr *ShopController) CreateShopWorkTime(ctx *gin.Context) {
	createShopWorkTimeRequest := shop_request.CreateShopWorkTimeRequest{}
	if err := ctx.ShouldBind(&createShopWorkTimeRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.CreateShopWorkTime(createShopWorkTimeRequest)
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

func (ctr *ShopController) CreateShopPhone(ctx *gin.Context) {
	createShopPhoneRequest := shop_request.CreateShopPhoneRequest{}
	if err := ctx.ShouldBind(&createShopPhoneRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.CreateShopPhone(createShopPhoneRequest)
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

func (ctr *ShopController) CreateShopSocial(ctx *gin.Context) {
	createShopSocialRequest := shop_request.CreateShopSocialRequest{}
	if err := ctx.ShouldBind(&createShopSocialRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.CreateShopSocial(createShopSocialRequest)
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}