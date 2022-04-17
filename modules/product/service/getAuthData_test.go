package service

import (
	"errors"

	"github.com/bagustyo92/auth/modules/auth/models"
	"github.com/golang/mock/gomock"
)

func (s *productServiceSuite) Test_productService_GetAuthData() {
	s.Run("#Case1: Err when get by token", func() {
		s.authService.EXPECT().GetByToken(gomock.Any()).Return(nil, errors.New("force err")).Times(1)

		data, err := s.productService.GetAuthData(models.JWTClaims{})
		s.NotNil(err)
		s.Equal(err, errors.New("force err"))
		s.Nil(data)
	})

	s.Run("#Case2: Success", func() {
		s.authService.EXPECT().GetByToken(gomock.Any()).Return(&models.Auth{}, nil).Times(1)

		data, err := s.productService.GetAuthData(models.JWTClaims{})
		s.Nil(err)
		s.NotNil(data)
		s.Equal(data, &models.Auth{})
	})
}
