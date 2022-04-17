package service

import (
	"testing"

	mock_auth_service "github.com/bagustyo92/auth/mocks/auth/service"
	mock_cache "github.com/bagustyo92/auth/mocks/cache"
	mock_request_converter "github.com/bagustyo92/auth/mocks/request/converter"
	mock_request_efishery "github.com/bagustyo92/auth/mocks/request/efishery"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type productServiceSuite struct {
	suite.Suite
	*require.Assertions

	ctrl *gomock.Controller

	authService      *mock_auth_service.MockAuthServiceInterface
	requestConverter *mock_request_converter.MockInterface
	requestEfishery  *mock_request_efishery.MockInterface
	cacher           *mock_cache.MockCacher

	productService ProductServiceInterface
}

func TestProductServiceSuite(t *testing.T) {
	suite.Run(t, new(productServiceSuite))
}

func (s *productServiceSuite) SetupTest() {
	s.Assertions = require.New(s.T())

	s.ctrl = gomock.NewController(s.T())

	s.authService = mock_auth_service.NewMockAuthServiceInterface(s.ctrl)
	s.requestConverter = mock_request_converter.NewMockInterface(s.ctrl)
	s.requestEfishery = mock_request_efishery.NewMockInterface(s.ctrl)
	s.cacher = mock_cache.NewMockCacher(s.ctrl)

	s.productService = NewService(s.authService, s.requestConverter, s.requestEfishery, s.cacher)
}

func (s *productServiceSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *productServiceSuite) TestNewService() {
	s.Run("#Case1", func() {
		svc := NewService(s.authService, s.requestConverter, s.requestEfishery, s.cacher)
		s.NotNil(svc)
	})
}
