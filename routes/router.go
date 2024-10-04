package router

import (
	"pennywise-api/routes/app"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo) {
	defaultGroup := e.Group("/v1")

	publicGroup := defaultGroup.Group("/auth")
	app.RegisterAuthRoutes(publicGroup.Group("/signup"))

	appGroup := defaultGroup.Group("/app")
	app.RegisterUserRoutes(appGroup.Group("/user"))
}
