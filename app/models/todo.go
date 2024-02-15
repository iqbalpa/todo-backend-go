package models

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title string `json:"title" binding:"required"`
	Description string `json:"desc" binding:"required"`
	isFinished bool `gorm:"default:false"`
}

// the isFinished will always set into false at initiation
func (todo *Todo) BeforeSave() error {
	if !todo.isFinished {
		todo.isFinished = false
	}
	return nil
}