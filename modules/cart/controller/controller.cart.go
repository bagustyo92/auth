package controller

import (
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/bagustyo92/auth/middleware/logger"
	"github.com/bagustyo92/auth/modules/cart/models"
	"github.com/bagustyo92/auth/utils"
	"github.com/labstack/echo"
)

func (uc *cartController) insertProduct(c echo.Context) error {
	cart := models.Cart{}
	if err := c.Bind(&cart); err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	res, err := uc.cartController.AddProductToCart(cart)
	if err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	return c.JSON(utils.Response(http.StatusOK, "OK", res))
}

func (uc *cartController) GetCart(c echo.Context) error {
	cartID := c.Param("id")
	filter := models.ProductCartFilter{}
	if err := c.Bind(&filter); err != nil {
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	id, err := uuid.FromString(cartID)
	if err != nil {
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	res, err := uc.cartController.GetCart(id, filter)
	if err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusInternalServerError, err, nil))
	}

	return c.JSON(utils.Response(http.StatusOK, "OK", res))
}

func (uc *cartController) deleteProductFromCart(c echo.Context) error {
	cartID := c.Param("cartID")
	cartIdUUID, err := uuid.FromString(cartID)
	if err != nil {
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}
	productCode := c.Param("productCode")
	if err := uc.cartController.DeleteProductFromCart(cartIdUUID, productCode); err != nil {
		logger.MakeLogEntry(c).Info(err)
		return c.JSON(utils.Response(http.StatusBadRequest, err, nil))
	}

	return c.JSON(utils.Response(http.StatusOK, "OK", nil))
}
