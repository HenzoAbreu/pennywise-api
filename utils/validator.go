package utils

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
)

type errorWrapper struct {
	Message string
}

func EmailValidator(c echo.Context, email string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if len(email) == 0 || len(email) > 255 || !re.MatchString(email) {
		return c.JSON(http.StatusBadRequest, errorWrapper{"invalid email field"})
	}
	return nil
}
func NameValidator(c echo.Context, name string) error {
	if len(name) == 0 || len(name) > 255 {
		return c.JSON(http.StatusBadRequest, errorWrapper{"invalid name field"})
	}
	return nil
}
func PhoneValidator(c echo.Context, phone string) error {
	re := regexp.MustCompile(`^\+?\d{7,15}$`)
	if !re.MatchString(phone) {
		return c.JSON(http.StatusBadRequest, errorWrapper{"invalid phone number"})
	}
	return nil
}
func CPFValidator(c echo.Context, cpf string) error {
	if len(cpf) == 0 || len(cpf) > 11 {
		return c.JSON(http.StatusBadRequest, errorWrapper{"invalid cpf field"})
	}
	return nil
}
func PasswordValidator(c echo.Context, password string) error {
	if len(password) < 6 || len(password) > 255 {
		return c.JSON(http.StatusBadRequest, errorWrapper{"invalid password field, password must have more than 6 characters"})
	}
	return nil
}
