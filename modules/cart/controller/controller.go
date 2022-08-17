package controller

import (
	"github.com/bagustyo92/auth/modules/cart/service"
	"github.com/labstack/echo"
)

type cartController struct {
	cartController service.CartInterface
}

func NewUserController(e *echo.Echo, cartService service.CartInterface) {
	handler := &cartController{
		cartController: cartService,
	}

	cart := e.Group("/cart")
	cart.GET("/:id", handler.GetCart)
	cart.POST("", handler.insertProduct)
	cart.DELETE("/:cartID/product/:productCode", handler.deleteProductFromCart)

	test := e.Group("/test")
	test.GET("/dockerfile", handler.dockerfile)
	test.GET("/money/:amount", handler.moneyTest)
	test.POST("/string", handler.stringTest)
	test.GET("/microservice", handler.microservice)
	test.GET("/dbindex", handler.indexDB)
}
