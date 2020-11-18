package controller

import (
	"fmt"
	"net/http"
	"web_app/logic"
	"web_app/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	// 1. verify parameter
	//var p models.ParamSignUp
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "params has error",
		})
		return
	}
	//fmt.Println(p)
	// verify

	if len(p.Username) == 0 || len(p.Password) == 0 || p.RePassword != p.Password {
		zap.L().Error("SignUp with invalid param")
		c.JSON(http.StatusOK, gin.H{
			"msg": "params has error",
		})
		return
	}
	fmt.Println(p)
	// 2. measurement
	logic.SignUp(p)
	// 3. return
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
