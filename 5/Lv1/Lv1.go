/*
* Lv1
* 写个登录注册接口，数据放到数据库中。
 */
package Lv1

import (
	"5/Lv1/router"
	"5/dao"
	"5/Lv1/model"
)

func Lv1()  {
	sql:=dao.Init(DBconf,&model.User{})
	server:=router.BuildRouter(sql)
	server.Run("localhost:8080")
}