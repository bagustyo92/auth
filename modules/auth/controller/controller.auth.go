package controller

import (
	"net/http"
	"time"

	"github.com/bagustyo92/auth/middleware/logger"
	"github.com/bagustyo92/auth/modules/auth/models"
	"github.com/bagustyo92/auth/utils"
	"github.com/labstack/echo"
)

func (ac *authController) login(c echo.Context) error {
	authString := c.Request().Header.Get("Authorization")

	user, err := ac.athController.Login(authString)
	if err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	token, err := ac.athController.CreateJwtToken(user)
	if err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusInternalServerError, err, nil))
	}

	auth := &models.Auth{
		UserID:      user.ID,
		AccessToken: token,
		LastLogin:   time.Now(),
	}
	if err := ac.athController.CreateAuth(auth); err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusInternalServerError, err, nil))
	}

	type response struct {
		User interface{}
		Auth interface{}
	}

	resp := response{
		user,
		auth,
	}

	return c.JSON(utils.Response(http.StatusOK, "OK", resp))
}
