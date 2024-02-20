package models

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title 		string `json:"title" binding:"required"`
	Description string `json:"desc" binding:"required"`
	IsFinished 	bool   `json:"isFinished" gorm:"default:false"`
	UserID		int    `json:"userId"`
}