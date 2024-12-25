package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	company_request "github.com/qahvazor/qahvazor/app/http/request/company"
	"github.com/qahvazor/qahvazor/app/http/response"
	"github.com/qahvazor/qahvazor/internal/service"
)

type CompanyController struct {
	service service.CompanyService
}

func NewCompanyController(service service.CompanyService) *CompanyController {
	return &CompanyController{
		service: service,
	}
}

func (ctr *CompanyController) GetCompanyList(ctx *gin.Context) {
	results, err := ctr.service.List()
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

func (ctr *CompanyController) GetCompanyShopsList(ctx *gin.Context) {
	getCompanyShopsRequest := company_request.GetCompanyShopsRequest{}
	if err := ctx.ShouldBind(&getCompanyShopsRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.GetCompanyShopsList(getCompanyShopsRequest)
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}
