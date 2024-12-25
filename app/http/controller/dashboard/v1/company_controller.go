package dashboard

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	company_request "github.com/qahvazor/qahvazor/app/http/request/company"
	company_social_request "github.com/qahvazor/qahvazor/app/http/request/company/social"
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

func (ctr *CompanyController) Store(ctx *gin.Context) {
	storeRequest := company_request.StoreRequest{}
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

func (ctr *CompanyController) Show(ctx *gin.Context) {
	companyParamID := ctx.Param("company_id")
	companyId, err := strconv.Atoi(companyParamID)
	if err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	result, err := ctr.service.Show(uint(companyId))
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

func (ctr *CompanyController) List(ctx *gin.Context) {
	results, err := ctr.service.List()
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

func (ctr *CompanyController) Edit(ctx *gin.Context) {
	editCompanyRequest := company_request.EditRequest{}
	if err := ctx.ShouldBind(&editCompanyRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}
	
	if err := ctr.service.Edit(editCompanyRequest); err != nil {
		response.NewResponse(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", nil, nil)
}

func (ctr *CompanyController) StoreCompanySocial(ctx *gin.Context) {
	storeCompanySocialRequest := company_social_request.StoreRequest{}
	if err := ctx.ShouldBind(&storeCompanySocialRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.StoreCompanySocial(storeCompanySocialRequest)
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	} 

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

func (ctr *CompanyController) ShowCompanySocial(ctx *gin.Context) {
	companySocialParamID := ctx.Param("social_id")
	socialId, err := strconv.Atoi(companySocialParamID)
	if err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	result, err := ctr.service.ShowCompanySocial(uint(socialId))
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

func (ctr *CompanyController) ListCompanySocials(ctx *gin.Context) {
	listCompanySocialRequest := company_social_request.ListRequest{}
	if err := ctx.ShouldBind(&listCompanySocialRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	results, err := ctr.service.ListCompanySocials(uint(listCompanySocialRequest.CompanyID))
	if err == nil {
		if errorResponse, ok := results.(response.ErrorResponse); ok {
			ctx.JSON(errorResponse.Code, errorResponse)
			return
		}
	} 

	response.NewResponse(ctx, http.StatusOK, "OK!", results, nil)
}

func (ctr *CompanyController) EditCompanySocial(ctx *gin.Context) {
	editCompanySocialRequest := company_social_request.EditRequest{}
	if err := ctx.ShouldBind(&editCompanySocialRequest); err != nil {
		response.NewResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}
	
	if err := ctr.service.EditCompanySocial(editCompanySocialRequest); err != nil {
		response.NewResponse(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	response.NewResponse(ctx, http.StatusOK, "OK!", nil, nil)
}
