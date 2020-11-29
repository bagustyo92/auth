package controller

import (
	"github.com/bagustyo92/auth/middleware/auth"
	"github.com/bagustyo92/auth/modules/user/service"
	"github.com/labstack/echo"
)

type userController struct {
	usrController service.UserServiceInterface
}

func NewUserController(e *echo.Echo, userService service.UserServiceInterface) {
	handler := &userController{
		usrController: userService,
	}

	user := e.Group("/user")
	user.POST("", handler.createUser)
	user.GET("/detail", handler.getUser, auth.JWTAuthorization)
	user.GET("", handler.getUsers, auth.JWTAuthorization)
	user.PATCH("/:id", handler.updateUser, auth.JWTAuthorization)
	user.DELETE("/:id", handler.deleteUser, auth.JWTAuthorization)
}
