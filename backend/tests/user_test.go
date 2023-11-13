package tests

import (
	"KitchenMistakeErazer/backend/handlers/users"
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestYourEchoHandler(t *testing.T) {
	// Створення екземпляра Echo
	e := echo.New()

	u := users.UsersHandler{}
	// Реєстрація вашого хендлера
	e.POST("/user/", u.SignUp)

	// Створення тестового запиту
	var jsonStr = []byte(`{
	"first_name": "Kseniia",
	"last_name": "Skopiuk",
	"email": "bg@example.com",
	"password": "dkjsfa"
}`)

	req := httptest.NewRequest(http.MethodPost, "/user/", bytes.NewBuffer(jsonStr))
	rec := httptest.NewRecorder()

	// Виклик вашого хендлера
	e.ServeHTTP(rec, req)

	// Перевірка статус коду та відповіді
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, this is your endpoint!", rec.Body.String())
}
