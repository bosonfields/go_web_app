package controller

import (
	"web_app/logic"
	"web_app/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func CreatePostHandler(c *gin.Context) {
	// 1. verify
	// 2. breate post
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShoudBindJSON(p) error", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	// get userID from c
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
	}
	p.AuthorId = userID
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3 return
	ResponseSuccess(c, nil)
}
