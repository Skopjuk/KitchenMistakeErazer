package users

import (
	"KitchenMistakeErazer/backend/repository"
	"KitchenMistakeErazer/backend/usecases/users"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

var (
	paginationLimit = "10"
)

type UserUpdateParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type PasswordUpdateAttributes struct {
	Password string `json:"password,omitempty"`
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

	usersRepository := repository.NewUsersRepository(u.container.DB)
	getAll := users.NewShowUsers(usersRepository)

	users, err := getAll.Execute(skip, paginationLimit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": users,
	})
}
