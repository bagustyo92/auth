package controller

import (
	"net/http"

	"github.com/bagustyo92/auth/middleware/logger"
	"github.com/bagustyo92/auth/modules/auth/models"
	"github.com/bagustyo92/auth/utils"
	"github.com/labstack/echo/v4"
)

// GetAccount godoc
// @Summary      Get account info using jwt token
// @Description  This endpoint will return account info based on token that received at this endpoint
// @Tags         Account
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Auth
// @Failure      400  {object}  utils.MOResponse
// @Failure      404  {object}  utils.MOResponse
// @Failure      500  {object}  utils.MOResponse
// @Router       /auth [get]
func (ac *authController) GetAuth(c echo.Context) error {
	authClaim := c.Get("auth").(models.JWTClaims)

	user, err := ac.athController.GetByToken(authClaim)
	if err != nil {
		logger.MakeLogEntry(c).Error(err)
		return c.JSON(utils.Response(http.StatusInternalServerError, err, nil))
	}
	return c.JSON(utils.Response(http.StatusOK, "OK", user))
}
