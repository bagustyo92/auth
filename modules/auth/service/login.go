package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/bagustyo92/auth/modules/auth/models"

	"github.com/dgrijalva/jwt-go"
)

var (
	CookieSessionLogin    = "SessionLogin"
	LoginSuccessCookieVal = "LoginSuccess"
	SecretKey             = "mLmHu8f1IxFo4dWurBG3jEf1Ex0wDZvvwND6eFmcaX"
	SigningMethod         = "HS512"
)

func (as *authService) Login(userLogin models.Auth) (string, error) {
	var (
		user *models.Auth
		err  error
	)

	// try login with phone number
	if userLogin.Phone != "" {
		if isPhoneNumber(userLogin.Phone) {
			return "", errors.New("Invalid Phone Number")
		}

		user, err = as.ar.GetByPhone(userLogin.Phone)
		if err != nil {
			return "", err
		}
	}

	// Try login using username
	if userLogin.Username != "" && user == nil {
		user, err = as.ar.GetByUsername(userLogin.Username)
		if err != nil {
			return "", errors.New("Username not found")
		}
	}

	if user == nil {
		return "", errors.New("Wrong username or password")
	}

	if userLogin.Password != user.Password || userLogin.Password == "" {
		return "", errors.New("wrong password")
	}

	return createJwtToken(user)
}

func createJwtToken(user *models.Auth) (string, error) {
	// we hash the jwt claims
	claims := models.JWTClaims{
		Name:      user.Username,
		Phone:     user.Phone,
		Role:      user.Role,
		CreatedAt: *user.CreatedAt,
		Id:        fmt.Sprintf("%d", user.ID),
		Issuer:    "efishery",
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(24)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
