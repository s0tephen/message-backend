package router

import (
	"MessageBoard/service/user_service"
	"github.com/gin-gonic/gin"
)

func Register(g *gin.Engine) {
	//创建路由组
	user := g.Group("/user")
	{
		user.POST("/sendMessage", user_service.SendMessage)
		user.GET("/showMessage", user_service.ShowMessage)
		user.POST("/detectMessage", user_service.DetectMessage)
	}
}
