package controller

import (
	"strconv"
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

func GetPostDetailHandler(c *gin.Context) {
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
	}

	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func GetPostListHandler(c *gin.Context) {

	page, size := getPageInfo(c)

	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func GetPostListHandler2(c *gin.Context) {
	// GET

	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}

	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetPostListNew(p)

	if err != nil {
		zap.L().Error("logic.GetPostList() failled", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}

	ResponseSuccess(c, data)

}

//func GetCommunityPostListHandler(c *gin.Context) {
//	p := &models.ParamCommunityPostList{
//		ParamPostList: &models.ParamPostList{
//			Page:  1,
//			Size:  10,
//			Order: models.OrderTime,
//		},
//	}
//
//	if err := c.ShouldBindQuery(p); err != nil {
//		zap.L().Error(" GetCommunityPostListHandler with invalid params", zap.Error(err))
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//
//	data, err := logic.GetCommunityPostList(p)
//
//	if err != nil {
//		zap.L().Error("logic.GetPostList() failled", zap.Error(err))
//		ResponseError(c, CodeServerBusy)
//	}
//
//	ResponseSuccess(c, data)
//}
