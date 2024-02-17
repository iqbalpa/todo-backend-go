package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username 	string `json:"username" binding:"required"`
	Password 	string `json:"password" binding:"required"`
	Name 		string `json:"name" binding:"required"`
}