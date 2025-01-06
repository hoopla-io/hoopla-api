package response

import (
	"github.com/gin-gonic/gin"
	"math"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
}

type Meta struct {
	ItemsPerPage int   `json:"itemsPerPage"`
	TotalItems   int64 `json:"totalItems"`
	CurrentPage  int   `json:"currentPage"`
	LastPage     int   `json:"lastPage"`
}

func (m *Meta) SetLastPage() {
	m.LastPage = int(math.Ceil(float64(m.TotalItems) / float64(m.ItemsPerPage)))
}

func NewResponse(c *gin.Context, statusCode int, message string, data interface{}, meta interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(statusCode, Response{
		Code:    statusCode,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func ValidationErrorResponse(c *gin.Context, message string) {
	NewResponse(c, 422, message, nil, nil)
}

func ServerErrorResponse(c *gin.Context, message string) {
	NewResponse(c, 500, message, nil, nil)
}

func BadRequestResponse(c *gin.Context, message string) {
	NewResponse(c, 400, message, nil, nil)
}

func NotFoundResponse(c *gin.Context, message string) {
	NewResponse(c, 404, message, nil, nil)
}

func InsufficientBalanceResponse(c *gin.Context, message string) {
	NewResponse(c, 428, message, nil, nil)
}

func ErrorResponse(c *gin.Context, code int, message string) {
	NewResponse(c, code, message, nil, nil)
}

func SuccessResponse(c *gin.Context, message string, data interface{}, meta interface{}) {
	NewResponse(c, 200, message, data, meta)
}
