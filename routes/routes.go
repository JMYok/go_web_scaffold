package routes

import (
	"github.com/gin-gonic/gin"
	"go_web_scaffold/logger"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})
	return r
}
