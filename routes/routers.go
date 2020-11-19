package routes

import (
	"net/http"
	"web_app/controller"
	"web_app/logger"
	"web_app/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//register route

	r.POST("/signup", controller.SignUpHandler)

	r.POST("/login", controller.LoginHandler)

	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		// if registered, judge where header exist valid JWT

		c.String(http.StatusOK, "pong")
		// please register
	})

	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "ok")
	//})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "404",
		})
	})

	return r
}
