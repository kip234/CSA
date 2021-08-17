/*
* Lv1
* 写个登录注册接口，数据放到数据库中。
 */
package Lv2

import (
	"5/Lv2/router"
	"5/dao"
	"5/Lv2/model"
)

func Lv2()  {
	sql:=dao.Init(DBconf,&model.Student{})
	server:=router.BuildRouter(sql)
	server.Run("localhost:8080")
}