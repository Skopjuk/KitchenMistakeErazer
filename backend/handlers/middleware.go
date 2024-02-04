package handlers

import (
	"KitchenMistakeErazer/backend/handlers/registration"
	"KitchenMistakeErazer/configs"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"strings"
)

func UserIdentityMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := userIdentity(c)
		if err != nil {
			return err
		}

		return next(c)
	}
}

func userIdentity(c echo.Context) error {
	header := c.Request().Header.Get("Authorization")
	if header == "" {
		logrus.Errorf("authorization header is empty, user: %s", c.Param("username"))
		return errors.New("authorization header is empty")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		logrus.Errorf("authorization header is invalid, user: %s", c.Param("username"))
		return errors.New("authorization header is invalid")
	}

	userId, _, err := parseToken(headerParts[1])
	if err != nil {
		logrus.Errorf("token can not be parsed")
		return errors.New(err.Error())
	}

	c.Set("userId", userId)

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
