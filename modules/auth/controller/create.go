package controller

import (
	"errors"
	"net/http"

	"github.com/bagustyo92/auth/middleware/logger"
	"github.com/bagustyo92/auth/modules/auth/models"
	"github.com/bagustyo92/auth/utils"
	"github.com/labstack/echo/v4"
)

// CreateAccount
// @Summary      Create or return existing account
// @Description  This endpoint will return account info or creating account if username is not exist, this endpoint also generating password for login
// @Tags         Account
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Auth
// @Failure      400  {object}  utils.MOResponse
// @Failure      404  {object}  utils.MOResponse
// @Failure      500  {object}  utils.MOResponse
// @Router       /auth/create [POST]
func (ac *authController) Create(c echo.Context) error {
	userCreate := models.Auth{}
	if err := c.Bind(&userCreate); err != nil {
		logger.MakeLogEntry(c).Error(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	if userCreate.Username == "" || userCreate.Phone == "" {
		return c.JSON(utils.Response(http.StatusBadRequest, errors.New("username or phone cannot be empty"), nil))
	}
	user, err := ac.athController.Create(userCreate)
	if err != nil {
		logger.MakeLogEntry(c).Error(err)
		return c.JSON(utils.Response(http.StatusInternalServerError, err, nil))
	}

	return c.JSON(utils.Response(http.StatusOK, "OK", user))
}
