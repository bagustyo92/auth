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

func (s *authControllerSuite) Test_authController_Create() {
	var (
		endpoint = "/auth/create"
	)

	s.Run("#Case1: Failed to bind", func() {
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
		_ = handler.Create(c)
		s.Equal(expectedStatusCode, rec.Code)
	})

	s.Run("#Case2: Username and phone empty", func() {
		requestBody, _ := json.Marshal(models.Auth{})

		e := echo.New()
		req := httptest.NewRequest(echo.POST, endpoint, bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		expectedStatusCode := http.StatusBadRequest

		handler := &authController{
			athController: s.authService,
		}
		_ = handler.Create(c)
		s.Equal(expectedStatusCode, rec.Code)
	})

	s.Run("#Case3: Service err", func() {
		requestBody, _ := json.Marshal(models.Auth{Username: "test"})

		s.authService.EXPECT().Create(gomock.Any()).Return(nil, errors.New("force err")).Times(1)

		e := echo.New()
		req := httptest.NewRequest(echo.POST, endpoint, bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		expectedStatusCode := http.StatusInternalServerError

		handler := &authController{
			athController: s.authService,
		}
		_ = handler.Create(c)
		s.Equal(expectedStatusCode, rec.Code)
	})

	s.Run("#Case4: Success", func() {
		requestBody, _ := json.Marshal(models.Auth{Username: "test"})

		s.authService.EXPECT().Create(gomock.Any()).Return(&models.Auth{}, nil).Times(1)

		e := echo.New()
		req := httptest.NewRequest(echo.POST, endpoint, bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		expectedStatusCode := http.StatusOK

		handler := &authController{
			athController: s.authService,
		}
		_ = handler.Create(c)
		s.Equal(expectedStatusCode, rec.Code)
	})
}
