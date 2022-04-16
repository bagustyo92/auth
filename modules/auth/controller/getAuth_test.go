package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/bagustyo92/auth/modules/auth/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func (s *authControllerSuite) Test_authController_GetAuth() {
	var (
		endpoint = "/auth"
	)

	s.Run("#Case1: Err when get by token", func() {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, endpoint, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		s.authService.EXPECT().GetByToken(gomock.Any()).Return(nil, errors.New("force err")).Times(1)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("auth", models.JWTClaims{})

		expectedStatusCode := http.StatusInternalServerError

		handler := &authController{
			athController: s.authService,
		}
		_ = handler.GetAuth(c)
		s.Equal(expectedStatusCode, rec.Code)
	})

	s.Run("#Case2: Success", func() {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, endpoint, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		s.authService.EXPECT().GetByToken(gomock.Any()).Return(&models.Auth{}, nil).Times(1)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("auth", models.JWTClaims{})

		expectedStatusCode := http.StatusOK

		handler := &authController{
			athController: s.authService,
		}
		_ = handler.GetAuth(c)
		s.Equal(expectedStatusCode, rec.Code)
	})
}
