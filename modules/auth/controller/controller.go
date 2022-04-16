package controller

import (
	midAuth "github.com/bagustyo92/auth/middleware/auth"
	"github.com/bagustyo92/auth/modules/auth/service"
	"github.com/labstack/echo/v4"
)

type authController struct {
	athController service.AuthServiceInterface
}

func NewAuthController(e *echo.Echo, authService service.AuthServiceInterface) {
	handler := &authController{
		athController: authService,
	}

	auth := e.Group("/auth")

	auth.POST("", handler.login)
	auth.POST("/create", handler.Create)
	auth.GET("", midAuth.JWTAuthorization(handler.GetAuth))
}
