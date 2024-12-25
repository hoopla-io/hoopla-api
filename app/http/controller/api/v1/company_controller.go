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

// @Tags Company
// @Accept  json
// @Produce  json
// @Success 200 {array} company_response.ListResponse "List of companies"
// @Router /company/list [get]
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

// @Tags Company
// @Accept  json
// @Produce  json
// @Param data body company_request.GetCompanyShopsRequest true "Get Company Shops Request"
// @Success 200 {array} company_response.GetCompanyShopsResponse "List of company shops"
// @Router /company/shops [post]
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
