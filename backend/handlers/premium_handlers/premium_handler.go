package premium_handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

func (p PremiumHandler) PremiumHandler(c echo.Context) error {
	token := c.Get("verifiedToken")
	if token == "" {
		http.Error(c.Response(), "Unauthorized", http.StatusUnauthorized)
		return errors.New("Unauthorized")
	}

	tmpl, err := template.ParseFiles("/Users/ksenia/KitchenMistakeErazer/static/premium.html")
	if err != nil {
		logrus.Printf("Error parsing templates: %v\n", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "this unit already exist",
		})
	}

	data := map[string]interface{}{
		"Token": token,
	}

	tmpl.Execute(c.Response().Writer, data)
	return nil
}
