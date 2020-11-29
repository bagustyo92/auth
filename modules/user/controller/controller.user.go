package controller

import (
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/bagustyo92/auth/middleware/logger"
	"github.com/bagustyo92/auth/modules/user/models"
	"github.com/bagustyo92/auth/utils"
	"github.com/labstack/echo"
)

func (uc *userController) createUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	if err := uc.usrController.CreateUser(&user); err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	return c.JSON(utils.Response(http.StatusOK, "OK", user))
}

func (uc *userController) getUser(c echo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}
	if err := uc.usrController.GetUser(user); err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}
	return c.JSON(utils.Response(http.StatusOK, "OK", user))
}

func (uc *userController) getUsers(c echo.Context) error {
	query := models.Query{}
	if err := c.Bind(&query); err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	res, err := uc.usrController.GetUsers(query)
	if err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	return c.JSON(utils.Response(http.StatusOK, "OK", res))
}

func (uc *userController) updateUser(c echo.Context) error {
	id := c.Param("id")
	userID, err := uuid.FromString(id)
	if err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}
	user := models.User{}
	user.ID = userID
	if err := c.Bind(&user); err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}
	if err := uc.usrController.UpdateUser(&user); err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}
	return c.JSON(utils.Response(http.StatusOK, "OK", user))
}

func (uc *userController) deleteUser(c echo.Context) error {
	id := c.Param("id")
	userID, err := uuid.FromString(id)
	if err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}
	if err := uc.usrController.DeleteUser(userID); err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}
	return c.JSON(utils.Response(http.StatusOK, "OK", "OK"))
}
