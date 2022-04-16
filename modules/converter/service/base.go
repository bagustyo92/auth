package service

import (
	"github.com/bagustyo92/auth/modules/auth/service"
	"github.com/bagustyo92/auth/modules/request/converter"
	"github.com/bagustyo92/auth/modules/request/efishery"
)

type converterService struct {
	authService       service.AuthServiceInterface
	currencyConverter converter.Interface
	efishery          efishery.Interface
}

func NewService(
	authSvc service.AuthServiceInterface,
	currConv converter.Interface,
	efishery efishery.Interface,
) ConverterServiceInterface {
	return &converterService{
		authService:       authSvc,
		currencyConverter: currConv,
		efishery:          efishery,
	}
}
