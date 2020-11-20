package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("user not login")

const CtxUserIDKey = "userID"

func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	//
	uid, ok := c.Get(CtxUserIDKey)
	if ok {
		userID, ok = uid.(int64)
	}
	if !ok {
		err = ErrorUserNotLogin
	}
	return
}
