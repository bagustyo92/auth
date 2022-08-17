package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bagustyo92/auth/config"
	"github.com/bagustyo92/auth/middleware/logger"
	"github.com/bagustyo92/auth/modules/cart/controller"
	"github.com/bagustyo92/auth/modules/cart/models"
	"github.com/bagustyo92/auth/modules/cart/repository"
	"github.com/bagustyo92/auth/modules/cart/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	envMode := os.Getenv("NODE_ENV")
	if envMode == "" {
		envMode = "development"
	}
	config.InitApp("env/" + envMode + ".env")

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", os.ModePerm)
	}

	var DBConnectionString string
	if config.DBUsername == "" || config.DBPassword == "" {
		DBConnectionString = fmt.Sprintf(":@tcp(%s:%s)/%s", config.DBURL, config.DBPort, config.DBName)
	} else {
		DBConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUsername, config.DBPassword, config.DBURL, config.DBPort, config.DBName)
	}
	db, err := gorm.Open("mysql", fmt.Sprintf("%s?charset=utf8&parseTime=True&loc=Local", DBConnectionString))

	// dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
	// 	config.DBURL, config.DBUsername, config.DBName, config.DBPassword)

	// gormClient, err := gorm.Open("postgres", dbUri)
	// if err != nil {
	// 	logger.MakeLogEntry(nil).Panic(fmt.Sprintf("create gorm client instance failed with message: %s", err.Error()))
	// 	os.Exit(1)
	// }
	db.LogMode(true)
	db.AutoMigrate(&models.Cart{}, &models.ProductCart{})

	defer db.Close()

	e := echo.New()
	e.Use(logger.Logging)
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		logger.MakeLogEntry(nil).Infof(strings.ReplaceAll(string(reqBody), "\r\n    ", ""))
	}))

	// Module Cart
	repo := repository.NewCartRepo(db)
	svc := service.NewCartService(repo)
	controller.NewUserController(e, svc)

	app := make(chan error)
	go func(app chan error) { app <- e.Start(":" + config.AppPort) }(app)
	err = <-app
	if err != nil {
		logger.MakeLogEntry(nil).Panic(err)
	}
}
