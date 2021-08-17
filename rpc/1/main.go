package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rpc/1/Models"
)

func main() {
	go server()
	s:=gin.Default()
	s.POST("/set", func(c *gin.Context){
		m:=Models.Message{}
		c.ShouldBind(&m)
		fmt.Println(m)
		_,re:=client(m,"KV.SetValue")
		c.JSON(http.StatusOK,gin.H{
			"message":re,
		})
	})
	s.GET("/get", func(c *gin.Context){
		m:=Models.Message{}
		c.ShouldBind(&m)
		fmt.Println(m)
		_,re:=client(m,"KV.GetValue")
		c.JSON(http.StatusOK,gin.H{
			"message":re,
		})
	})
	s.Run(":8080")
}
