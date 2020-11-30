package controller

import (
	"strconv"
	"web_app/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetEventHandler(c *gin.Context) {

	//eventName := c.Param("name")
	//pid, err := strconv.ParseInt(pidStr, 10, 64)
	//if err != nil {
	//	zap.L().Error("get post detail with invalid param", zap.Error(err))
	//}

	eidStr := c.Param("id")
	eid, err := strconv.ParseInt(eidStr, 10, 64)
	if err != nil {
		zap.L().Error("get event detail with invalid param", zap.Error(err))
	}

	data, err := logic.GetEventById(eid)
	if err != nil {
		zap.L().Error("logic.GetEventById(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
