package dto

import (
	"github.com/jinzhu/gorm"
)

type LoginRequest struct {
	gorm.Model
	Username 	string `json:"username"`
	Password 	string `json:"password"`
}