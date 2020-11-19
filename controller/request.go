package controller

import (
	"errors"
	"web_app/middlewares"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("user not login")

func getCurrentUser(c *gin.Context) (userID int64, err error) {
	//
	uid, ok := c.Get(middlewares.CtxUserIDKey)
	if ok {
		userID, ok = uid.(int64)
	}
	if !ok {
		err = ErrorUserNotLogin
	}
	return
}
