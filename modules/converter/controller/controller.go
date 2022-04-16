package controller

import (
	midAuth "github.com/bagustyo92/auth/middleware/auth"
	"github.com/bagustyo92/auth/modules/converter/service"
	"github.com/labstack/echo/v4"
)

type converterController struct {
	converterService service.ConverterServiceInterface
}

func NewController(e *echo.Echo, convCtrl service.ConverterServiceInterface) {
	handler := &converterController{
		converterService: convCtrl,
	}

	products := e.Group("/product-list")

	products.GET("", midAuth.JWTAuthorization(handler.GetProductList))
	products.GET("/summary", midAuth.JWTAuthorization(handler.GetSummaryPrice))
	products.GET("/data-user", midAuth.JWTAuthorization(handler.GetAuthData))
}
