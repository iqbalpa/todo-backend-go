package utils

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=iqbalpahlevi dbname=todo_go port=5432 sslmode=disable"

	var err error
	DB, err = gorm.Open("postgres", dsn)
	
	if err != nil {
		fmt.Println("Failed to connect database ", err.Error())
	} else {
		fmt.Println("Connected to database postgres")
	}
}
