package service

import (
	"errors"
	"time"

	"github.com/bagustyo92/auth/modules/auth/models"
	"github.com/golang/mock/gomock"
)

func (s *authServiceSuite) Test_authService_Login() {
	var (
		forceErr = errors.New("force err")
	)

	s.Run("#Case1: Err invalid phone number", func() {
		token, err := s.authService.Login(models.Auth{Phone: "test09928928"})
		s.NotNil(err)
		s.Empty(token)
		s.Equal(err, errors.New("Invalid Phone Number"))
	})

	s.Run("#Case2: Failed get by phone", func() {
		s.authRepo.EXPECT().GetByPhone(gomock.Any()).Return(nil, forceErr).Times(1)

		token, err := s.authService.Login(models.Auth{Phone: "082277738847"})

		s.NotNil(err)
		s.Empty(token)
		s.Equal(err, forceErr)
	})

	s.Run("#Case3: Failed get by username", func() {
		s.authRepo.EXPECT().GetByUsername(gomock.Any()).Return(nil, forceErr).Times(1)

		token, err := s.authService.Login(models.Auth{Username: "test09928928"})

		s.NotNil(err)
		s.Empty(token)
		s.Equal(err, errors.New("Username not found"))
	})

	s.Run("#Case4: Err wrong username or password", func() {
		token, err := s.authService.Login(models.Auth{})
		s.NotNil(err)
		s.Empty(token)
		s.Equal(err, errors.New("Wrong username or password"))
	})

	s.Run("#Case5: Wrong password", func() {
		s.authRepo.EXPECT().GetByUsername(gomock.Any()).Return(&models.Auth{}, nil).Times(1)

		token, err := s.authService.Login(models.Auth{Password: "tsjsfakjnsf", Username: "test"})

		s.NotNil(err)
		s.Empty(token)
		s.Equal(err, errors.New("wrong password"))
	})

	s.Run("#Case6: Success", func() {
		s.authRepo.EXPECT().GetByUsername(gomock.Any()).Return(&models.Auth{Password: "test", Base: models.Base{CreatedAt: &time.Time{}}}, nil).Times(1)

		token, err := s.authService.Login(models.Auth{Password: "test", Username: "test"})

		s.Nil(err)
		s.NotEmpty(token)
	})
}
