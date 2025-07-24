package vendor_controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hoopla/hoopla-api/app/http/response"
	"github.com/hoopla/hoopla-api/internal/service"
	"net/http"
)

type IikoController struct {
	partnerService   service.PartnerService
	userOrderService service.UserOrderService
}

func NewIikoController(userOrderService service.UserOrderService, partnerService service.PartnerService) *IikoController {
	return &IikoController{
		partnerService:   partnerService,
		userOrderService: userOrderService,
	}
}

// @Tags IIKO
// @Accept  json
// @Produce  json
// @Param data body vendors_poster_request.WebhookRequest true "Webhook for iiko"
// @Router /vendors/iiko/webhook [post]
func (c *IikoController) Webhook(ctx *gin.Context) {
	var request []map[string]interface{}
	if err := ctx.ShouldBind(&request); err != nil {
		response.ValidationErrorResponse(ctx, err.Error())
		return
	}

	fmt.Println("---------------")
	fmt.Println("request", request)
	fmt.Println("---------------")

	if request[0]["eventType"] == "DeliveryOrderUpdate" {
		eventInfo := request[0]["eventInfo"].(map[string]interface{})

		partner, code, err := c.partnerService.GetPartnerByVendorId(eventInfo["organizationId"].(string))
		if err != nil {
			if code == http.StatusNotFound {
				response.NotFoundResponse(ctx, eventInfo["organizationId"].(string))
				return
			}
			response.ErrorResponse(ctx, code, err.Error())
			return
		}

		userOrder, code, err := c.userOrderService.GetOrderByVendorOrderID(
			partner.ID,
			partner.Vendor,
			eventInfo["id"].(string),
		)
		if err != nil {
			response.ErrorResponse(ctx, code, err.Error())
			return
		}

		//in case if order was missing
		userOrder.PartnerID = partner.ID
		userOrder.Vendor = partner.Vendor
		userOrder.VendorOrderID = eventInfo["id"].(string)

		orderStatus := "pending"
		order := eventInfo["order"].(map[string]interface{})
		if order["status"].(string) == "CookingStarted" {
			orderStatus = "preparing"
		} else if order["status"].(string) == "CookingCompleted" {
			orderStatus = "completed"
		} else if order["status"].(string) == "Waiting" {
			orderStatus = "completed"
		} else if order["status"].(string) == "Delivered" {
			orderStatus = "completed"
		} else if order["status"].(string) == "Closed" {
			orderStatus = "completed"
		} else if order["status"].(string) == "Cancelled" {
			orderStatus = "cancelled"
		}

		userOrder, code, err = c.userOrderService.UpdateOrderStatus(userOrder, orderStatus)
		if err != nil {
			response.ErrorResponse(ctx, code, err.Error())
			return
		}

		response.SuccessResponse(ctx, "OK!", orderStatus, nil)
		return
	}

	response.SuccessResponse(ctx, "OK!", nil, nil)
}
