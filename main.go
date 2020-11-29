package main

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	authService "github.com/bagustyo92/auth/modules/auth/service"
	"github.com/bagustyo92/auth/modules/user/service"

	"github.com/bagustyo92/auth/config"

	"github.com/bagustyo92/auth/middleware/logger"
	authController "github.com/bagustyo92/auth/modules/auth/controller"
	authModels "github.com/bagustyo92/auth/modules/auth/models"
	authRepo "github.com/bagustyo92/auth/modules/auth/repository"
	"github.com/bagustyo92/auth/modules/user/controller"
	"github.com/bagustyo92/auth/modules/user/models"
	"github.com/bagustyo92/auth/modules/user/repository"
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
	db.AutoMigrate(&models.User{}, &authModels.Auth{})

	defer db.Close()

	e := echo.New()
	e.Use(logger.Logging)
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		logger.MakeLogEntry(nil).Infof(strings.ReplaceAll(string(reqBody), "\r\n    ", ""))
	}))

	// Module User
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	controller.NewUserController(e, userService)

	// Module Auth
	authRepo := authRepo.NewAuthRepo(db)
	auth := authService.NewAuthService(authRepo, userRepo)
	authController.NewAuthController(e, auth)

	app := make(chan error)
	go func(app chan error) { app <- e.Start(":" + config.AppPort) }(app)
	err = <-app
	if err != nil {
		logger.MakeLogEntry(nil).Panic(err)
	}
}
