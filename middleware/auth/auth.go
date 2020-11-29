package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/bagustyo92/auth/modules/auth/models"

	"github.com/bagustyo92/auth/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
)

var JWTSigningMethod = jwt.SigningMethodHS256

// JWTSignature Must Provide in Cred File
var JWTSignature = []byte("mLmHu8f1IxFo4dWurBG3jEf1Ex0wDZvvwND6eFmcaX")

// JWTAuthorization is middleware to parse token from request header
func JWTAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var tokenString string
		authHeader := c.Request().Header.Get("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			return c.JSON(utils.Response(http.StatusUnauthorized, "Invalid Token", nil))
		}
		tokenString = strings.Replace(authHeader, "Bearer ", "", -1)

		// Token parse and validate time
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Signing method invalid")
			} else if method != JWTSigningMethod {
				return nil, errors.New("Signing method invalid")
			}
			return JWTSignature, nil
		})

		if err != nil {
			return c.JSON(utils.Response(http.StatusUnauthorized, "Unauthorized Token", nil))
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.JSON(utils.Response(http.StatusUnauthorized, "Token Has Expired", nil))
		}

		var jwtModel models.JWTClaims
		mapstructure.Decode(claims, &jwtModel)
		// res, _ := json.Marshal(jwtModel)

		/*
			>>> Unit Test Purposes <<<
			- Uncomment this line if you want to pass the unit test
		*/
		// c.Response().Header().Set("JWT_PARSE", string(res))

		// Add claims to context
		c.Set("auth", jwtModel)
		return next(c)
	}
}
