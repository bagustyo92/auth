package service

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bagustyo92/auth/modules/auth/models"

	userModels "github.com/bagustyo92/auth/modules/user/models"
	"github.com/bagustyo92/auth/utils"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

var (
	CookieSessionLogin    = "SessionLogin"
	LoginSuccessCookieVal = "LoginSuccess"
	SecretKey             = "mLmHu8f1IxFo4dWurBG3jEf1Ex0wDZvvwND6eFmcaX"
	SigningMethod         = "HS512"
)

func (as *authService) basicParse(authString string) (string, string, error) {
	authSplit := strings.SplitN(authString, " ", 2)

	if len(authSplit) != 2 || authSplit[0] != "Basic" {
		return "", "", fmt.Errorf("Invalid authorization string")
	}

	decoded, _ := base64.StdEncoding.DecodeString(authSplit[1])

	var identifier, password string
	authPair := strings.SplitN(string(decoded), ":", 2)
	if len(authPair) < 2 {
		identifier = authPair[0]
		password = ""
	} else {
		identifier = authPair[0]
		password = authPair[1]
	}

	return identifier, password, nil
}

func (as *authService) Login(authString string) (*userModels.User, error) {
	identifier, password, err := as.basicParse(authString)
	if err != nil {
		return nil, err
	}

	fmt.Println(identifier, password)
	user := userModels.User{
		UserName: identifier,
	}
	// user.UserName = identifier

	fmt.Println(user)
	if err := as.ur.GetUser(&user); err != nil {
		return nil, errors.New("username not found")
	}

	if err := utils.ComparePwd(user.Password, password); err != nil {
		return nil, errors.New("wrong password")
	}

	// Set Cookies
	return &user, nil
}

func (as *authService) CreateJwtToken(user *userModels.User) (string, error) {
	// we hash the jwt claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID.String(),
		"iss":  "lemonilo",
		"app":  "app.lemonilo.com",
		"role": "user",
		"name": user.Name,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * time.Duration(24)).Unix(),
	})

	accessToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (as *authService) CreateAuth(auth *models.Auth) error {
	return as.ar.CreateAuth(auth)
}

func (as *authService) UpdateAuth(auth *models.Auth) error {
	return as.ar.UpdateAuth(auth)
}

func (as *authService) GetAuth(userID uuid.UUID) (*models.Auth, error) {
	return as.ar.GetAuth(userID)
}
