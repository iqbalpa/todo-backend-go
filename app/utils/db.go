package utils

import (
	"fmt"
	"main/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	config := config.LoadEnv()
	dsn := config.GetDSN()

	var err error
	DB, err = gorm.Open(config.DatabaseDriver, dsn)
	
	if err != nil {
		fmt.Println("Failed to connect database ", err.Error())
	} else {
		fmt.Println("Connected to database postgres")
	}
}
