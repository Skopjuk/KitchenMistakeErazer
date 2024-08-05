package handlers

import (
	"KitchenMistakeErazer/backend/handlers/registration"
	"KitchenMistakeErazer/configs"
	"context"
	"errors"
	"firebase.google.com/go/v4/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

//func UserIdentityMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		err := userIdentity(c)
//		if err != nil {
//			return err
//		}
//
//		return next(c)
//	}
//}

func userIdentity(c echo.Context, client auth.Client) error {
	params, err := url.ParseQuery(c.Request().URL.RawQuery)
	if err != nil {
		logrus.Errorf("query can not be parsed")
		return errors.New(err.Error())
	}

	idToken := params.Get("auth-token")
	if idToken == "" {
		logrus.Errorf("idToken can not be empty")
		return errors.New(err.Error())
	}

	// verify ID token
	verifiedToken, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		logrus.Printf("Error verifying token: %v\n", err)
		return errors.New(err.Error())
	}

	// Add token to the request context
	c.Set("verifiedToken", verifiedToken)
	return nil
}

func parseToken(accessToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &registration.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		config, err := configs.NewConfig()
		if err != nil {
			logrus.Error("config is not available")
			return "", err
		}

		return []byte(config.SigningKey), nil
	})

	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*registration.TokenClaims)
	if !ok {
		return 0, "", errors.New("token claims are not of type *TokenClaims")
	}

	return claims.UserId, claims.Role, nil
}

func AuthMiddleWare(client *auth.Client) func(echo2 echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			// parse Query parameters
			params, err := url.ParseQuery(c.Request().URL.RawQuery)
			if err != nil {
				http.Error(c.Response(), "Bad Request", http.StatusBadRequest)
				// log.Fatalf("")
				return errors.New("Bad Request")
			}

			idToken := params.Get("auth-token")
			if idToken == "" {
				http.Error(c.Response(), "Unauthorized", http.StatusUnauthorized)
				return errors.New("Unauthorized")
			}

			// verify ID token
			verifiedToken, err := client.VerifyIDToken(context.Background(), idToken)
			if err != nil {
				http.Error(c.Response(), "Unauthorized", http.StatusUnauthorized)
				logrus.Printf("Error verifying token: %v\n", err)
			}

			// Add token to the request context
			c.Set("verifiedToken", verifiedToken)
			return nil
		})
	}
}
