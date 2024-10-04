package app

import (
	"pennywise-api/controller"
	"pennywise-api/controller/viewmodel"
	"pennywise-api/routes/routeutils"

	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(e *echo.Group) {
	const (
		createUser = ""
	)

	e.POST(createUser, handleCreateUser)
}

func handleCreateUser(c echo.Context) error {
	request := viewmodel.SignupRequest{}
	c.Bind(&request)
	user, err := controller.CreateUser(c, request)
	if err != nil {
		return err
	}

	return routeutils.ResponseAPIOK(c, user)
}
