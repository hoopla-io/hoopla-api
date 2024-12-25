package dashboard

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	coffee_request "github.com/qahvazor/qahvazor/app/http/request/coffee"
	"github.com/qahvazor/qahvazor/app/http/response"
	"github.com/qahvazor/qahvazor/internal/service"
)

type CoffeeController struct {
	service service.CoffeeService
}

func NewCoffeeController(service service.CoffeeService) *CoffeeController {
	return &CoffeeController{
		service: service,
	}
}

func (ctr *CoffeeController) Store(ctx *gin.Context) {
	storeRequest := coffee_request.StoreRequest{}
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

func (ctr *CoffeeController) Show(ctx *gin.Context) {
	coffeeParamID := ctx.Param("coffee_id")
	coffeeId, err := strconv.Atoi(coffeeParamID)
	if err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	result, err := ctr.service.Show(uint(coffeeId))
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

func (ctr *CoffeeController) List(ctx *gin.Context) {
	results, err := ctr.service.List()
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

func (ctr *CoffeeController) Edit(ctx *gin.Context) {
	editCoffeeRequest := coffee_request.EditRequest{}
	if err := ctx.ShouldBind(&editCoffeeRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}
	
	if err := ctr.service.Edit(editCoffeeRequest); err != nil {
		response.NewResponse(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", nil, nil)
}