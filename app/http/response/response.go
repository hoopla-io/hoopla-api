package response

import (
	"github.com/gin-gonic/gin"
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

func NewResponse(c *gin.Context, statusCode int, message string, data interface{}, meta interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(statusCode, Response{
		Code:    statusCode,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}