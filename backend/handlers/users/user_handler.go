package users

import (
	"KitchenMistakeErazer/backend/repository/user"
	"KitchenMistakeErazer/backend/usecases/users"
	"fmt"
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

type UserUpdateParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
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

func (u *UsersHandler) UpdateUser(c echo.Context) (err error) {
	var input users.UpdateUserAttributes

	idInt, err := getIdFromEndpoint(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("users id %d can not be parsed", idInt),
		})
	}

	if err := c.Bind(&input); err != nil {
		logrus.Errorf("failed to bind req body: %s", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	usersRepository := user.NewUsersRepository(u.container.DB)
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

func (u *UsersHandler) DeleteUser(c echo.Context) error {
	idInt, err := getIdFromEndpoint(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("users id %d can not be parsed", idInt),
		})
	}

	usersRepository := user.NewUsersRepository(u.container.DB)
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

func getIdFromEndpoint(c echo.Context) (int, error) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Errorf("error of converting id to int. id: %s", id)
		return 0, c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": fmt.Sprintf("id %d can not be parsed", idInt),
		})
	}

	return idInt, nil
}
