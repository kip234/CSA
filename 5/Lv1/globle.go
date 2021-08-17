package Lv1

import (
	"5/dao"
	"gorm.io/gorm"
)
//数据库配置
var DBconf=dao.Sql{
	"homework",
	"root",
	"root",
	"127.0.0.1:3306",
}

var DB *gorm.DB