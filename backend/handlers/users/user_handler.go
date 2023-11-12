package users

import (
	"KitchenMistakeErazer/backend/repository/user"
	"KitchenMistakeErazer/backend/usecases/users"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

var (
	paginationLimit = "10"
)

type SignUpParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u *UsersHandler) SignUp(c echo.Context) error {
	var input SignUpParams //треба спробуавти одразу розпарсити в users.UserAttributes
	if err := c.Bind(&input); err != nil {
		logrus.Errorf("error while binding sign up input: %s", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	usersRepository := user.NewUsersRepository(u.container.DB)
	createProfile := users.NewCreateUserProfile(usersRepository)

	userInsertingParams := users.UserAttributes{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	}

	err := createProfile.Execute(userInsertingParams)
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
		"added_user": input,
	})
	if err != nil {
		logrus.Errorf("troubles with sending http status: %s", err)
	}

	return nil

}

func (u *UsersHandler) GetAllUsers(c echo.Context) error {
	pageNum := c.QueryParam("page")
	if pageNum == "" {
		pageNum = "1"
	}

	page, err := strconv.Atoi(pageNum)
	if err != nil {
		logrus.Errorf("error while converting page number to int: %s", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "error while parsing url",
		})
	}

	skip := strconv.Itoa((page - 1) * 10)
	logrus.Info("attempt to get users list from db")

	usersRepository := user.NewUsersRepository(u.container.DB)
	getAll := users.NewShowUsers(usersRepository)

	users, err := getAll.Execute(skip, paginationLimit)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": users,
	})
}
