package controller

import (
	"encoding/json"
	"testing"

	mock_service "github.com/bagustyo92/auth/mocks/product/service"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type productControllerSuite struct {
	suite.Suite
	*require.Assertions

	ctrl           *gomock.Controller
	productService *mock_service.MockProductServiceInterface
}

func (s *productControllerSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())
	s.productService = mock_service.NewMockProductServiceInterface(s.ctrl)
}

func TestAuthController(t *testing.T) {
	suite.Run(t, new(productControllerSuite))
}

func (s *productControllerSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *productControllerSuite) MockExpectedResp(data interface{}) string {
	expectedByte, _ := json.Marshal(data)
	return string(expectedByte) + "\n"
}

func (s *productControllerSuite) TestNewController() {
	s.Run("#Case1", func() {
		NewController(echo.New(), s.productService)
	})
}
