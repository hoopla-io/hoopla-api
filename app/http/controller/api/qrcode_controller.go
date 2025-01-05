package api

import (
	"github.com/gin-gonic/gin"
	"github.com/qahvazor/qahvazor/internal/service"
	"net/http"
)

type QRCodeController struct {
	QRCodeService *service.QRCodeService
}

func NewQRCodeController(qrCodeService *service.QRCodeService) *QRCodeController {
	return &QRCodeController{
		QRCodeService: qrCodeService,
	}
}

func (c *QRCodeController) GenerateQRCode(ctx *gin.Context) {
	data := ctx.Query("data")
	if data == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "data query parameter is required"})
		return
	}

	qrCode := c.QRCodeService.GenerateQRCode(data)
	ctx.JSON(http.StatusOK, gin.H{"qr_code": qrCode})
}

func (c *QRCodeController) DecodeQRCode(ctx *gin.Context) {
	encodedData := ctx.Query("encoded_data")
	if encodedData == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "encoded_data query parameter is required"})
		return
	}

	data, err := c.QRCodeService.DecodeQRCode(encodedData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": data})
}
