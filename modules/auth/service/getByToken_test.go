package service

import (
	"errors"

	"github.com/bagustyo92/auth/modules/auth/models"
	"github.com/golang/mock/gomock"
)

func (s *authServiceSuite) Test_authService_GetByToken() {
	var (
		forceErr = errors.New("force err")
	)

	s.Run("#Case1: Unknown err", func() {
		s.authRepo.EXPECT().GetByUsername(gomock.Any()).Return(nil, forceErr).Times(1)
		data, err := s.authService.GetByToken(models.JWTClaims{})
		s.Nil(data)
		s.NotNil(err)
		s.Equal(err, forceErr)
	})

	s.Run("#Case2: Success", func() {
		s.authRepo.EXPECT().GetByUsername(gomock.Any()).Return(&models.Auth{}, nil).Times(1)
		data, err := s.authService.GetByToken(models.JWTClaims{})
		s.Nil(err)
		s.NotNil(data)
		s.Equal(data, &models.Auth{})
	})
}
