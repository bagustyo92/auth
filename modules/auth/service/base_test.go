package service

import (
	"testing"

	mock_repository "github.com/bagustyo92/auth/mocks/auth/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type authServiceSuite struct {
	suite.Suite
	*require.Assertions

	ctrl *gomock.Controller

	authRepo    *mock_repository.MockAuthRepoInterface
	authService AuthServiceInterface
}

func TestAuthServiceSuite(t *testing.T) {
	suite.Run(t, new(authServiceSuite))
}

func (s *authServiceSuite) SetupTest() {
	s.Assertions = require.New(s.T())

	s.ctrl = gomock.NewController(s.T())

	s.authRepo = mock_repository.NewMockAuthRepoInterface(s.ctrl)
	s.authService = NewAuthService(s.authRepo)
}

func (s *authServiceSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *authServiceSuite) TestNewAuthService() {
	s.Run("#Case1", func() {
		svc := NewAuthService(s.authRepo)
		s.NotNil(svc)
	})
}
