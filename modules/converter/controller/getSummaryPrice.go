package controller

import (
	"errors"
	"net/http"

	"github.com/bagustyo92/auth/middleware/logger"
	"github.com/bagustyo92/auth/modules/auth/models"
	"github.com/bagustyo92/auth/utils"
	"github.com/labstack/echo/v4"
)

// GetSummaryProductPrice from products godoc
// @Summary      This endpoint will return summary price based on data that fetched from efishery, based on province and also time
// @Description  This endpoint will return summary price based on data that fetched from efishery, based on province and also time
// @Tags         Product
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.SummaryPriceList
// @Failure      400  {object}  utils.MOResponse
// @Failure      404  {object}  utils.MOResponse
// @Failure      500  {object}  utils.MOResponse
// @Router       /product-list/summary [get]
func (cc *converterController) GetSummaryPrice(c echo.Context) error {
	authClaim := c.Get("auth").(models.JWTClaims)

	if authClaim.Role == "" {
		return c.JSON(utils.Response(http.StatusUnauthorized, errors.New("You dont have any roles!"), nil))
	}

	if authClaim.Role != "admin" {
		return c.JSON(utils.Response(http.StatusUnauthorized, errors.New("You dont have permission!"), nil))
	}

	data, err := cc.converterService.GetSummaryPriceList()
	if err != nil {
		logger.MakeLogEntry(c).Error(err)
		return c.JSON(utils.Response(http.StatusInternalServerError, err, nil))
	}

	return c.JSON(utils.Response(http.StatusOK, "OK", data))
}
