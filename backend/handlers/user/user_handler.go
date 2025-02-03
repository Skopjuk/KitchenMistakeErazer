package user

import (
	"KitchenMistakeErazer/backend/handlers"
	"KitchenMistakeErazer/backend/repository"
	"KitchenMistakeErazer/backend/usecases/users"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
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

func (u *UserHandler) GetUserByID(c echo.Context) error {
	idInt, err := handlers.GetIdFromEndpoint(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("users id %d can not be parsed", idInt),
		})
	}

	usersRepository := repository.NewUsersRepository(u.container.DB)
	getAll := users.NewGetUserByID(usersRepository)

	user, err := getAll.Execute(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func (u *UserHandler) UpdateUser(c echo.Context) (err error) {
	var input users.UpdateUserAttributes

	idInt, err := handlers.GetIdFromEndpoint(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("users id %d can not be parsed", idInt),
		})
	}

	if err := c.Bind(&input); err != nil {
		logrus.Errorf("failed to bind req body: %s", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	usersRepository := repository.NewUsersRepository(u.container.DB)
	newGetUserById := users.NewGetUserByID(usersRepository)
	result, err := newGetUserById.Execute(idInt)
	if err != nil {
		logrus.Errorf("problem while extracting user from DB: %s", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("problem while extracting user from DB: %s", err),
		})
	}

	if result.FirstName == "" {
		logrus.Errorf(fmt.Sprintf("user with id %d wasn't found", idInt))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": fmt.Sprintf("user with id %d wasn't found", idInt),
		})
	}

	newUpdateProfile := users.NewUpdateUserProfile(usersRepository)
	err = newUpdateProfile.Execute(input, idInt)
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"user": input.Email,
	})
	if err != nil {
		logrus.Errorf("troubles with sending http status: %s", err)
	}

	return err
}

func (u *UserHandler) DeleteUser(c echo.Context) error {
	idInt, err := handlers.GetIdFromEndpoint(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("users id %d can not be parsed", idInt),
		})
	}

	usersRepository := repository.NewUsersRepository(u.container.DB)
	newGetUserById := users.NewGetUserByID(usersRepository)
	_, err = newGetUserById.Execute(idInt)
	if err != nil {
		logrus.Errorf("problem while extracting user from DB: %s", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("problem while extracting user from DB: %s", err),
		})
	}

	newDeleteUser := users.NewDeleteUserProfile(usersRepository)
	err = newDeleteUser.Execute(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("user deletion useccessful: %s", err),
		})
	}

	return err
}

func (u *UserHandler) UpdateUsersPassword(c echo.Context) error {
	var input PasswordUpdateAttributes
	idInt, err := handlers.GetIdFromEndpoint(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("users id %d can not be parsed", idInt),
		})
	}

	usersRepository := repository.NewUsersRepository(u.container.DB)
	newGetUserById := users.NewGetUserByID(usersRepository)
	_, err = newGetUserById.Execute(idInt)
	if err != nil {
		logrus.Errorf("problem while extracting user from DB: %s", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("problem while extracting user from DB: %s", err),
		})
	}

	if err := c.Bind(&input); err != nil {
		logrus.Errorf("failed to bind req body: %s", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	newUpdateUsersPassword := users.NewChangePassword(usersRepository)
	err = newUpdateUsersPassword.Execute(idInt, input.Password)
	if err != nil {
		logrus.Errorf("user password wasn't updated: %s", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("user password wasn't updated: %s", err),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "password changed successfully",
	})
}
