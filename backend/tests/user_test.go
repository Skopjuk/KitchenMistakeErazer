package tests

import (
	"KitchenMistakeErazer/backend/container"
	"KitchenMistakeErazer/backend/handlers/users"
	"KitchenMistakeErazer/configs"
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestYourEchoHandler(t *testing.T) {
	// Створення екземпляра Echo
	e := echo.New()
	logging := logrus.New()
	logging.SetReportCaller(true)

	// Імітуємо спосіб завантаження конфігурації в реальному додатку
	config, err := loadConfig()
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	containerInstance := container.NewContainer(config, logging)

	u := users.NewUsersHandler(containerInstance)
	// Реєстрація вашого хендлера
	e.POST("/user/", u.SignUp)

	// Створення тестового запиту
	var jsonStr = []byte(`{
        "first_name": "Ksdeniia",
        "last_name": "Skodpiuk",
        "email": "bsssg@example.com",
        "password": "dkjsfddda"
    }`)

	req := httptest.NewRequest(http.MethodPost, "/user/", bytes.NewBuffer(jsonStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Виклик вашого хендлера
	e.ServeHTTP(rec, req)

	// Перевірка статус коду та відповіді
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, this is your endpoint!", rec.Body.String())
}

// Метод для завантаження конфігурації, імітуючи дійсний метод
func loadConfig() (configs.Config, error) {
	// Ваша логіка завантаження конфігурації тут
	return configs.NewConfig()
}
