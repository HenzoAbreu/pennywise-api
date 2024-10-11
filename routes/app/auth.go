package app

import (
	"pennywise-api/controller"
	"pennywise-api/controller/viewmodel"
	"pennywise-api/routes/routeutils"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func RegisterAuthRoutes(e *echo.Group) {
	const (
		createUser = ""
	)

	e.POST(createUser, handleCreateUser)
}

func handleCreateUser(c echo.Context) (err error) {
	request := viewmodel.SignupRequest{}

	if err = c.Bind(&request); err != nil {
		return errors.Wrap(err, "failed to bind request")
	}

	user, err := controller.CreateUser(c, request)
	if err != nil {
		return errors.Wrap(err, "failed to create user")
	}

	return routeutils.ResponseCreated(c, user)
}
