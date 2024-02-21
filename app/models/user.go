package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username 	string `json:"username" binding:"required" gorm:"unique"`
	Password 	string `json:"password" binding:"required"`
	Name 		string `json:"name" binding:"required"`
	Todos		[]*Todo `json:"todos" gorm:"foreignKey:UserID"`
}