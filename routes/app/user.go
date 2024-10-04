package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Group) {
	const (
		loggedUserRoute = ""
	)
	e.GET(loggedUserRoute, handleGetUser)
}

func handleGetUser(c echo.Context) error {
	return c.String(http.StatusOK, "Got the usersss")
}
