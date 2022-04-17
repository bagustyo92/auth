package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/bagustyo92/auth/modules/auth/models"
	reqModel "github.com/bagustyo92/auth/modules/request/models"
	"github.com/labstack/echo/v4"
)

func (s *productControllerSuite) Test_productController_GetProductList() {
	var (
		endpoint = "/product-list"
	)

	s.Run("#Case1: Err when role is empty", func() {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, endpoint, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("auth", models.JWTClaims{
			Role: "",
		})

		expectedStatusCode := http.StatusUnauthorized

		handler := &productController{
			converterService: s.productService,
		}
		_ = handler.GetProductList(c)
		s.Equal(expectedStatusCode, rec.Code)
	})

	s.Run("#Case2: Err when get price list", func() {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, endpoint, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		s.productService.EXPECT().GetPriceListIncludingPriceInUSD().Return(nil, errors.New("force err")).Times(1)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("auth", models.JWTClaims{
			Role: "admin",
		})

		expectedStatusCode := http.StatusInternalServerError

		handler := &productController{
			converterService: s.productService,
		}
		_ = handler.GetProductList(c)
		s.Equal(expectedStatusCode, rec.Code)
	})

	s.Run("#Case3: Success get product", func() {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, endpoint, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		s.productService.EXPECT().GetPriceListIncludingPriceInUSD().Return([]reqModel.Price{}, nil).Times(1)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("auth", models.JWTClaims{
			Role: "admin",
		})

		expectedStatusCode := http.StatusOK

		handler := &productController{
			converterService: s.productService,
		}
		_ = handler.GetProductList(c)
		s.Equal(expectedStatusCode, rec.Code)
	})
}
