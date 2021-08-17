package handler

import (
	"5/Lv2/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Stu(DB *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		students:=model.Load(DB)
		//message,err:=json.Marshal(students)
		//if err!=nil{
			//c.JSON(http.StatusOK,gin.H{
			//	"error":err.Error(),
			//})
		//}else{
		c.JSON(http.StatusOK,&students)
		//}
	}
}
