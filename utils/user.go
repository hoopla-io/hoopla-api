package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

type UserHelper struct {
	UserID      uint
	PhoneNumber string
}

func (u *UserHelper) Init(ctx *gin.Context) error {
	authorization := strings.Split(ctx.GetHeader("Authorization"), " ")
	if len(authorization) == 2 {
		token := authorization[1]
		claims, err := DecodeJWT(token)
		if err != nil {
			return err
		}

		u.UserID = claims.UserID
		u.PhoneNumber = claims.PhoneNumber
		return nil
	}

	return errors.New("authorization header is not valid")
}
