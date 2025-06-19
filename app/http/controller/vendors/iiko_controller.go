package vendor_controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hoopla/hoopla-api/app/http/response"
	"github.com/hoopla/hoopla-api/internal/service"
)

type IikoController struct {
	userOrderService    service.UserOrderService
	partnerTokenService service.PartnerTokenService
}

func NewIikoController(userOrderService service.UserOrderService, partnerTokenService service.PartnerTokenService) *IikoController {
	return &IikoController{
		userOrderService:    userOrderService,
		partnerTokenService: partnerTokenService,
	}
}

// @Tags IIKO
// @Accept  json
// @Produce  json
// @Param data body vendors_poster_request.WebhookRequest true "Webhook for iiko"
// @Router /vendors/iiko/webhook [post]
func (c *IikoController) Webhook(ctx *gin.Context) {
	var r []map[string]interface{}
	if err := ctx.ShouldBind(&r); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	fmt.Println("IIIKOOO__------------------------")
	fmt.Println(r)
	fmt.Println("IIIKOOO__------------------------")

	response.SuccessResponse(ctx, "OK!", nil, nil)
}
