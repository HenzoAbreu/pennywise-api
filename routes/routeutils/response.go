package routeutils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type responseWrapper struct {
	Error   bool
	Message string
	Data    interface{}
}

func ResponseAPIOK(c echo.Context, data interface{}) error {
	retVal := responseWrapper{
		Error:   false,
		Message: "success",
		Data:    data,
	}

	return c.JSON(http.StatusOK, retVal)
}
