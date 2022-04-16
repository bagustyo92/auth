package controller

import (
	"encoding/json"
	"testing"

	mock_service "github.com/bagustyo92/auth/mocks/auth/service"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type authControllerSuite struct {
	suite.Suite
	*require.Assertions

	ctrl *gomock.Controller

	authService *mock_service.MockAuthServiceInterface
}

func (s *authControllerSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())
	s.authService = mock_service.NewMockAuthServiceInterface(s.ctrl)
}

func TestAuthController(t *testing.T) {
	suite.Run(t, new(authControllerSuite))
}

func (s *authControllerSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (cs *authControllerSuite) MockExpectedResp(data interface{}) string {
	expectedByte, _ := json.Marshal(data)
	return string(expectedByte) + "\n"
}

func (s *authControllerSuite) TestNewAuthController() {
	s.Run("#Case1", func() {
		NewAuthController(echo.New(), s.authService)
	})
}
