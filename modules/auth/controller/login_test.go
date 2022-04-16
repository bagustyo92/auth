package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/bagustyo92/auth/modules/auth/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func (s *authControllerSuite) Test_authController_login() {
	var (
		endpoint = "/auth"
	)

	s.Run("#Case1: Err failed to bind", func() {
		requestBody, _ := json.Marshal("test")

		e := echo.New()
		req := httptest.NewRequest(echo.POST, endpoint, bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		expectedStatusCode := http.StatusBadRequest

		handler := &authController{
			athController: s.authService,
		}
		_ = handler.login(c)
		s.Equal(expectedStatusCode, rec.Code)
	})

	s.Run("#Case2: Err login", func() {
		requestBody, _ := json.Marshal(models.Auth{})

		e := echo.New()
		req := httptest.NewRequest(echo.POST, endpoint, bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		s.authService.EXPECT().Login(gomock.Any()).Return("", errors.New("force err")).Times(1)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		expectedStatusCode := http.StatusInternalServerError

		handler := &authController{
			athController: s.authService,
		}
		_ = handler.login(c)
		s.Equal(expectedStatusCode, rec.Code)
	})

	s.Run("#Case3: Success Login and get token", func() {
		requestBody, _ := json.Marshal(models.Auth{})

		e := echo.New()
		req := httptest.NewRequest(echo.POST, endpoint, bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		s.authService.EXPECT().Login(gomock.Any()).Return("token", nil).Times(1)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		expectedStatusCode := http.StatusOK

		handler := &authController{
			athController: s.authService,
		}
		_ = handler.login(c)
		s.Equal(expectedStatusCode, rec.Code)
	})
}
