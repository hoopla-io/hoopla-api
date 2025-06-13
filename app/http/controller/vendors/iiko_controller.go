package vendor_controllers

import "github.com/gin-gonic/gin"

type IikoController struct{}

func NewIikoController() *IikoController {
	return &IikoController{}
}

func (c *IikoController) Webhook(ctx *gin.Context) {

}
