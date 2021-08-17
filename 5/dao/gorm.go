package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(sql Sql,model interface{}) *gorm.DB {
	dsn := sql.SqlUserName+":"+sql.SqlUserPwd+"@tcp("+sql.SqlAddr+")/"+sql.SqlName+"?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(model)
	return db
}