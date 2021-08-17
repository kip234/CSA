package model

import (
	"gorm.io/gorm"
	"fmt"
	)

type User struct {
	Id uint `form:"id"`
	Pwd string `form:"pwd"`
	Name string `form:"name"`
}

func (u *User)Is(db *gorm.DB) (err error) {
	if u.Id==0{
		err=fmt.Errorf("u.Id==0")
		return
	}
	tmp:=User{}
	err=db.Where("id=?",u.Id).Find(&tmp).Error
	if err!=nil{
		return
	}
	fmt.Println(tmp.Pwd, u.Pwd)
	if tmp.Pwd!=u.Pwd{
		err=fmt.Errorf("pwd wrong")
	}
	return
}