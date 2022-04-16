package service

import (
	"errors"

	"github.com/bagustyo92/auth/modules/auth/models"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
)

func (s *authServiceSuite) Test_authService_Create() {
	var (
		forceErr = errors.New("force err")
	)

	s.Run("#Case1: Unknown err", func() {
		s.authRepo.EXPECT().GetByUsername(gomock.Any()).Return(nil, forceErr).Times(1)
		data, err := s.authService.Create(models.Auth{})
		s.Nil(data)
		s.NotNil(err)
		s.Equal(err, forceErr)
	})

	s.Run("#Case2: Err record not found and err when create", func() {
		s.authRepo.EXPECT().GetByUsername(gomock.Any()).Return(nil, gorm.ErrRecordNotFound).Times(1)
		s.authRepo.EXPECT().Create(gomock.Any()).Return(nil, forceErr).Times(1)

		data, err := s.authService.Create(models.Auth{})
		s.Nil(data)
		s.NotNil(err)
		s.Equal(err, forceErr)
	})

	s.Run("#Case3: Success Create", func() {
		s.authRepo.EXPECT().GetByUsername(gomock.Any()).Return(nil, gorm.ErrRecordNotFound).Times(1)
		s.authRepo.EXPECT().Create(gomock.Any()).Return(&models.Auth{}, nil).Times(1)

		data, err := s.authService.Create(models.Auth{})
		s.Nil(err)
		s.NotNil(data)
		s.Equal(data, &models.Auth{})
	})
}
