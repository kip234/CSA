package router

import (
	"5/Lv2/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BuildRouter(DB *gorm.DB) *gin.Engine {
	server:=gin.Default()
	group:=server.Group("/")
	{
		group.POST("/login",handler.Login(DB))//只建立了一条路由，因为题目需要
		group.GET("/stu",handler.Stu(DB))
	}
	return server
}


