package service

import (
	"pennywise-api/entity"
	"pennywise-api/utils"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context, user entity.User) (retVal entity.User, err error) {
	// full validation process
	err = utils.EmailValidator(c, user.Email)
	if err != nil {
		return retVal, err
	}
	err = utils.NameValidator(c, user.FullName)
	if err != nil {
		return retVal, err
	}
	err = utils.PhoneValidator(c, user.Phone)
	if err != nil {
		return retVal, err
	}
	err = utils.CPFValidator(c, user.CPF)
	if err != nil {
		return retVal, err
	}
	err = utils.PasswordValidator(c, user.Password)
	if err != nil {
		return retVal, err
	}
	return user, err
}
