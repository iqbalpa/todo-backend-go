package app

import (
	"log"
	"main/app/models"
	"main/app/routes"
	"main/app/utils"
	"main/config"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type App struct {
	api    *gin.RouterGroup
	db     *gorm.DB
	router *gin.Engine
}

func NewApp() *App {
	// connect to database
	utils.ConnectDB()
	// migrate models
	utils.DB.AutoMigrate(&models.Todo{})

	// initiate router
	router := gin.Default()
	router.Use(gin.Recovery())

	// grouping api
	api := router.Group("/api/v1")

	return &App{
		api:    api,
		db:     utils.DB,
		router: router,
	}
}

func (a *App) Run() {
	config := config.LoadEnv()

	api := a.api
	db := a.db

	serverPort := config.GetServerPort()
	routes.TodoRoutes(api, db)

	a.router.Run(serverPort)
}