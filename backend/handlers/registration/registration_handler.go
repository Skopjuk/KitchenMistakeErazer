package registration

import (
	"KitchenMistakeErazer/backend/repository"
	"KitchenMistakeErazer/backend/usecases/users"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type SignUpParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
}

type TokenClaims struct {
	jwt.RegisteredClaims
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}

func (u *SignUpHandler) SignUp(c echo.Context) error {
	var input SignUpParams //треба спробуавти одразу розпарсити в users.UserAttributes
	if err := c.Bind(&input); err != nil {
		logrus.Errorf("error while binding sign up input: %s", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	usersRepository := repository.NewUsersRepository(u.container.DB)
	createProfile := users.NewCreateUserProfile(usersRepository)

	userInsertingParams := users.UserAttributes{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	}

	id, err := createProfile.Execute(userInsertingParams)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "this user already exist",
			})
		}

		logrus.Errorf("error while executing inserting user params: %s", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		logrus.Errorf("troubles with sending http status: %s", err)
	}

	return nil
}
