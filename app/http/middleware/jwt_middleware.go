package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/qahvazor/qahvazor/app/http/response"
	"github.com/qahvazor/qahvazor/utils"
	"net/http"
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
