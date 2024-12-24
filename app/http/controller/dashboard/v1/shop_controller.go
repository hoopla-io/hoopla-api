package dashboard

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	shop_request "github.com/qahvazor/qahvazor/app/http/request/shop"
	shop_phone_request "github.com/qahvazor/qahvazor/app/http/request/shop/phone"
	shop_worktime_request "github.com/qahvazor/qahvazor/app/http/request/shop/worktime"
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

func (ctr *ShopController) Store(ctx *gin.Context) {
	storeRequest := shop_request.StoreRequest{}
	if err := ctx.ShouldBind(&storeRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.Store(storeRequest)
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

func (ctr *ShopController) Show(ctx *gin.Context) {
	shopParamID := ctx.Param("shop_id")
	shopId, err := strconv.Atoi(shopParamID)
	if err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	result, err := ctr.service.Show(uint(shopId))
	if err != nil {
		response.NewResponse(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	response.NewResponse(
		ctx,
		http.StatusOK,
		"OK!",
		result,
		nil,
	)
}

func (ctr *ShopController) List(ctx *gin.Context) {
	results, err := ctr.service.List()
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

func (ctr *ShopController) Edit(ctx *gin.Context) {
	editRequest := shop_request.EditRequest{}
	if err := ctx.ShouldBind(&editRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}
	
	if err := ctr.service.Edit(editRequest); err != nil {
		response.NewResponse(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", nil, nil)
}


func (ctr *ShopController) StoreShopWorktime(ctx *gin.Context) {
	storeShopWorktimeRequest := shop_worktime_request.StoreRequest{}
	if err := ctx.ShouldBind(&storeShopWorktimeRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.StoreShopWorktime(storeShopWorktimeRequest)
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	} 

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

func (ctr *ShopController) ShowWorktime(ctx *gin.Context) {
	shopWorktimeParamID := ctx.Param("worktime_id")
	worktimeId, err := strconv.Atoi(shopWorktimeParamID)
	if err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	result, err := ctr.service.ShowWorktime(uint(worktimeId))
	if err != nil {
		response.NewResponse(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	response.NewResponse(
		ctx,
		http.StatusOK,
		"OK!",
		result,
		nil,
	)
}

func (ctr *ShopController) ListShopWorktimes(ctx *gin.Context) {
	listShopWorktimeRequest := shop_worktime_request.ListRequest{}
	if err := ctx.ShouldBind(&listShopWorktimeRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.ListShopWorktimes(uint(listShopWorktimeRequest.ShopID))
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	} 

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

func (ctr *ShopController) EditShopWorktime(ctx *gin.Context) {
	editShopWorktimeRequest := shop_worktime_request.EditRequest{}
	if err := ctx.ShouldBind(&editShopWorktimeRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}
	
	if err := ctr.service.EditShopWorktime(editShopWorktimeRequest); err != nil {
		response.NewResponse(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", nil, nil)
}

func (ctr *ShopController) StoreShopPhone(ctx *gin.Context) {
	storeShopPhoneRequest := shop_phone_request.StoreRequest{}
	if err := ctx.ShouldBind(&storeShopPhoneRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.StoreShopPhone(storeShopPhoneRequest)
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	} 

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

func (ctr *ShopController) ShowShopPhone(ctx *gin.Context) {
	shopPhoneParamID := ctx.Param("phone_id")
	phoneId, err := strconv.Atoi(shopPhoneParamID)
	if err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	result, err := ctr.service.ShowShopPhone(uint(phoneId))
	if err != nil {
		response.NewResponse(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	response.NewResponse(
		ctx,
		http.StatusOK,
		"OK!",
		result,
		nil,
	)
}

func (ctr *ShopController) ListShopPhones(ctx *gin.Context) {
	listShopPhoneRequest := shop_phone_request.ListRequest{}
	if err := ctx.ShouldBind(&listShopPhoneRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.ListShopPhones(uint(listShopPhoneRequest.PhoneID))
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	} 

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

func (ctr *ShopController) EditShopPhone(ctx *gin.Context) {
	editShopPhoneRequest := shop_phone_request.EditRequest{}
	if err := ctx.ShouldBind(&editShopPhoneRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}
	
	if err := ctr.service.EditShopPhone(editShopPhoneRequest); err != nil {
		response.NewResponse(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", nil, nil)
}
