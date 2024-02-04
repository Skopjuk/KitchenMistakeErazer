package login

import (
	"KitchenMistakeErazer/backend/container"
	"KitchenMistakeErazer/backend/handlers/registration"
	"KitchenMistakeErazer/backend/repository"
	"KitchenMistakeErazer/backend/usecases/users"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type SignInParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const timeToLimit = 1

func (u *SignInHandler) SignIn(c echo.Context) error {
	var input SignInParams
	logrus.Infof("user %s tries to authenticate", input)

	if err := c.Bind(&input); err != nil {
		u.container.Logging.Errorf("failed to bind req body: %s", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	params := users.AuthenticationAttributes{
		Email:    input.Email,
		Password: input.Password,
	}

	token, err := GenerateToken(params.Email, params.Password, *u.container)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "user do not exist",
		})
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
	if err != nil {
		logrus.Errorf("troubles with sending http status: %s", err)
		err = c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"token": token,
		})
	}

	return err
}

func GenerateToken(email, password string, c container.Container) (string, error) {
	params := users.AuthenticationAttributes{
		Email:    email,
		Password: password,
	}
	usersRepository := repository.NewUsersRepository(c.DB)
	newAuthentication := users.NewAuthentication(usersRepository)
	foundUser, err := newAuthentication.Execute(params)
	if err != nil {
		logrus.Errorf("cannot execute usecase: %s", err.Error())
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &registration.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * timeToLimit)),
		},
		UserId: foundUser.Id,
	})

	signedString, err := token.SignedString([]byte(c.Config.SigningKey))
	if err != nil {
		return "", err
	}

	return signedString, nil
}
