package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name string `gorm:"type:varchar(20);not null" form:"name"`
	Pwd string `gorm:"size:255;not null" form:"pwd"`
}

//type UserLogin struct {
//	Username string `json:"username"`
//	Password string `json:"password"`
//	Code string `json:"code"`
//}