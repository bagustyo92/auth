package main

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/bagustyo92/auth/modules/auth/models"
	"github.com/bagustyo92/auth/modules/request/converter"
	"github.com/bagustyo92/auth/modules/request/efishery"

	"github.com/bagustyo92/auth/config"

	authCtrl "github.com/bagustyo92/auth/modules/auth/controller"
	authRepository "github.com/bagustyo92/auth/modules/auth/repository"
	authSvc "github.com/bagustyo92/auth/modules/auth/service"

	productCtrl "github.com/bagustyo92/auth/modules/converter/controller"
	productSvc "github.com/bagustyo92/auth/modules/converter/service"

	"github.com/bagustyo92/auth/middleware/logger"
	"github.com/labstack/echo/v4"

	_ "github.com/bagustyo92/auth/docs"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Efishery Test API
// @version 1.0
// @description This is a api based on efishery test purposes.
// @termsOfService http://swagger.io/terms/

// @contact.name Bagus Prasetyo
// @contact.url https://github.com/bagustyo92
// @contact.email bagustyo92@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8999
// @BasePath /
func main() {
	envMode := os.Getenv("NODE_ENV")
	if envMode == "" {
		envMode = "development"
	}
	config.InitApp("env/" + envMode + ".env")

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", os.ModePerm)
	}

	fmt.Println(config.DBURL, config.DBName, config.DBPassword, config.DBPort)

	var DBConnectionString string
	if config.DBUsername == "" || config.DBPassword == "" {
		DBConnectionString = fmt.Sprintf(":@tcp(%s:%s)/%s", config.DBURL, config.DBPort, config.DBName)
	} else {
		DBConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUsername, config.DBPassword, config.DBURL, config.DBPort, config.DBName)
	}
	dsn := fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local", DBConnectionString)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&models.Auth{})

	defer db.Close()

	e := echo.New()
	e.Use(logger.Logging)
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		logger.MakeLogEntry(nil).Infof(strings.ReplaceAll(string(reqBody), "\r\n    ", ""))
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Request().URL.Path, "swagger") {
				return true
			}
			return false
		},
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Module Auth
	auth := authRepository.NewRepository(db)
	authService := authSvc.NewAuthService(auth)
	authCtrl.NewAuthController(e, authService)

	// Module Product
	productService := productSvc.NewService(authService, converter.NewRequest(), efishery.NewRequest())
	productCtrl.NewController(e, productService)

	app := make(chan error)
	go func(app chan error) { app <- e.Start(":" + config.AppPort) }(app)
	err = <-app
	if err != nil {
		logger.MakeLogEntry(nil).Panic(err)
	}
}
