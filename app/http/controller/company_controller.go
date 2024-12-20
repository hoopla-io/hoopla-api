package controller

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

func (ctr *CompanyController) CreateCompany(ctx *gin.Context) {
	createCompanyRequest := company_request.CreateCompanyRequest{}
	if err := ctx.ShouldBind(&createCompanyRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.CreateCompany(createCompanyRequest)
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

func (ctr *CompanyController) GetCompany(ctx *gin.Context) {
	getCompanyRequest := company_request.GetCompanyRequest{}
	if err := ctx.ShouldBindQuery(&getCompanyRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.GetCompany(getCompanyRequest)
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

func (ctr *CompanyController) GetList(ctx *gin.Context) {
	results, err := ctr.service.GetList()
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}