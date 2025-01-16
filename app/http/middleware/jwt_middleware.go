package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoopla/hoopla-api/app/http/response"
	"github.com/hoopla/hoopla-api/utils"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userHelper utils.UserHelper
		err := userHelper.Init(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, response.Response{
				Code:    http.StatusPreconditionFailed,
				Message: err.Error(),
				Data:    nil,
				Meta:    nil,
			})
		}

		ctx.Next()
	}
}
