package service

import (
	"pennywise-api/data"
	"pennywise-api/entity"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func CreateUser(c echo.Context, user entity.User) (retVal entity.User, err error) {
	// trim spaces

	// create salt and uuid

	// check if user email or cpf or phone number are already registered

	user.PasswordSalt = "14078956498576234592387"
	user.UUID = "123456789012345678901234567890123456"

	user, err = data.Save(user)
	if err != nil {
		return retVal, errors.Wrap(err, "failed to create user")
	}

	return user, err
}
