package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/bagustyo92/auth/modules/auth/models"
	prodModel "github.com/bagustyo92/auth/modules/product/models"
	"github.com/labstack/echo/v4"
)

func (s *productControllerSuite) Test_productController_GetSummaryPrice() {
	var (
		endpoint = "/product-list/summary"
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
		_ = handler.GetSummaryPrice(c)
		s.Equal(expectedStatusCode, rec.Code)
	})

	s.Run("#Case2: Err when not admin", func() {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, endpoint, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("auth", models.JWTClaims{
			Role: "user",
		})

		expectedStatusCode := http.StatusUnauthorized

		handler := &productController{
			converterService: s.productService,
		}
		_ = handler.GetSummaryPrice(c)
		s.Equal(expectedStatusCode, rec.Code)
	})

	s.Run("#Case3: Failed get summary price list", func() {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, endpoint, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("auth", models.JWTClaims{
			Role: "admin",
		})

		expectedStatusCode := http.StatusInternalServerError

		s.productService.EXPECT().GetSummaryPriceList().Return(nil, errors.New("force err")).Times(1)

		handler := &productController{
			converterService: s.productService,
		}
		_ = handler.GetSummaryPrice(c)
		s.Equal(expectedStatusCode, rec.Code)
	})

	s.Run("#Case4: Failed get summary price list", func() {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, endpoint, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("auth", models.JWTClaims{
			Role: "admin",
		})

		expectedStatusCode := http.StatusOK

		s.productService.EXPECT().GetSummaryPriceList().Return(&prodModel.SummaryPriceList{}, nil).Times(1)

		handler := &productController{
			converterService: s.productService,
		}
		_ = handler.GetSummaryPrice(c)
		s.Equal(expectedStatusCode, rec.Code)
	})
}
