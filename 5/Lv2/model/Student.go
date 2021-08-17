package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Student struct {
	Name string `json:"name" form:"name" gorm:"name"`
	StuNum string `json:"stunum" form:"stunum" gorm:"stunum"`//学号
	Class string `json:"class" form:"class" gorm:"class"`
	Major string `json:"major" form:"major" gorm:"major"`
}

func (s *Student)Is(db *gorm.DB) (err error) {
	var tmp Student
	err=db.Where("name=?",s.Name).Find(&tmp).Error
	if err!=nil{
		return
	}
	if tmp.StuNum!=s.StuNum{
		fmt.Println(tmp.StuNum)
		fmt.Println(s.StuNum)
		err=fmt.Errorf("????>")
	}
	return
}

func Load(DB *gorm.DB) (member Student) {
	//tmp:=Student{}
	DB.Where("name=?","kip").Find(&member)
	return
}