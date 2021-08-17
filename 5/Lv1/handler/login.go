package handler

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"5/Lv1/model"
	"net/http"
	)

//登录
func Login(DB *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		err:=c.ShouldBind(&user)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{"Error":err.Error()})
		}
		if err=user.Is(DB);err!=nil{
			c.JSON(http.StatusOK,gin.H{"Error":err.Error()})
		}else{
			c.JSON(http.StatusOK,user)
			//c.JSON(http.StatusOK,gin.H{
			//	"message":"welcome",
			//	"id":user.Id,
			//	"name":user.Name,
			//	"pwd":user.Pwd,
			//
			//})
		}
	}
}
