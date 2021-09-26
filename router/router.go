package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xueqiya/go_project/controller"
)

func Setup() *gin.Engine {
	r := gin.New()
	// 中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apiv1 := r.Group("/api/v1")
	{
		// ding 路由
		apiv1.POST("/ding/sendSms", controller.SendSms)
	}
	return r
}