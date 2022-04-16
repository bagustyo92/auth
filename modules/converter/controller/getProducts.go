package controller

import (
	"errors"
	"net/http"

	"github.com/bagustyo92/auth/middleware/logger"
	"github.com/bagustyo92/auth/modules/auth/models"
	reqModel "github.com/bagustyo92/auth/modules/request/models"
	"github.com/bagustyo92/auth/utils"
	"github.com/labstack/echo/v4"
)

// GetProductList from products godoc
// @Summary      Fetch product from efishery endpoint and add data price in usd at response
// @Description  Fetch product from efishery endpoint and add data price in usd at response
// @Tags         Product
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.MOResponse
// @Failure      400  {object}  utils.MOResponse
// @Failure      404  {object}  utils.MOResponse
// @Failure      500  {object}  utils.MOResponse
// @Router       /product-list [get]
func (cc *converterController) GetProductList(c echo.Context) error {
	var (
		err  error
		data []reqModel.Price
	)
	authClaim := c.Get("auth").(models.JWTClaims)

	if authClaim.Role == "" {
		return c.JSON(utils.Response(http.StatusUnauthorized, errors.New("You dont have any roles"), nil))
	}

	data, err = cc.converterService.GetPriceListIncludingPriceInUSD()
	if err != nil {
		logger.MakeLogEntry(c).Error(err)
		return c.JSON(utils.Response(http.StatusInternalServerError, err, nil))
	}

	return c.JSON(utils.Response(http.StatusOK, "OK", data))
}
