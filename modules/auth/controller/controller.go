package controller

import (
	"github.com/bagustyo92/auth/modules/auth/service"
	"github.com/labstack/echo"
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
}
