package router

import (
	"MessageBoard/service/user_service"
	"github.com/gin-gonic/gin"
)

func Register(g *gin.Engine) {
	//创建路由组
	user := g.Group("/user")
	{
		user.GET("/showMessage", user_service.ShowMessage)
		user.POST("/detectMessage", user_service.DetectMessage)
		user.POST("/sendMessageAndUpdate", user_service.SendMessageAndUpdate)
	}
}
