package main

import (
	"pennywise-api/config"
	database "pennywise-api/db"
	router "pennywise-api/routes"
	"pennywise-api/routes/routeutils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//Echo server instance

	e := echo.New()

	// Load config
	config.LoadConfig()

	// Connect to the database
	database.ConnectDatabase()

	// middlware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//custom error handler
	e.HTTPErrorHandler = routeutils.CustomHTTPErrorHandler

	//get the routes
	router.Register(e)

	//start server
	appAdress := "localhost:1323"
	e.Logger.Fatal(e.Start(appAdress))
}
