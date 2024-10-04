package main

import (
	router "pennywise-api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//Echo server instance

	e := echo.New()

	// middlware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//get the routes
	router.Register(e)

	//start server
	appAdress := "localhost:1323"
	e.Logger.Fatal(e.Start(appAdress))
}
