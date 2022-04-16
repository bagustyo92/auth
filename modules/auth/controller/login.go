package controller

import (
	"net/http"

	"github.com/bagustyo92/auth/middleware/logger"
	"github.com/bagustyo92/auth/modules/auth/models"
	"github.com/bagustyo92/auth/utils"
	"github.com/labstack/echo/v4"
)

// Login godoc
// @Summary      This endpoint receive username, phone and password for getting the JWT token
// @Description  This endpoint will return jwt token based on username/phone and password that received at this endpoint
// @Tags         Account
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.LoginResp
// @Failure      400  {object}  utils.MOResponse
// @Failure      404  {object}  utils.MOResponse
// @Failure      500  {object}  utils.MOResponse
// @Router       /auth [post]
func (ac *authController) login(c echo.Context) error {
	userLogin := models.Auth{}
	if err := c.Bind(&userLogin); err != nil {
		logger.MakeLogEntry(c).Error(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}
	token, err := ac.athController.Login(userLogin)
	if err != nil {
		logger.MakeLogEntry(c).Error(err)
		return c.JSON(utils.Response(http.StatusInternalServerError, err, nil))
	}

	return c.JSON(utils.Response(http.StatusOK, "OK", models.LoginResp{AccessToken: token}))
}
