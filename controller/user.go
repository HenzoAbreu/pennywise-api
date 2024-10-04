package controller

import (
	"pennywise-api/controller/viewmodel"
	"pennywise-api/service"

	"github.com/labstack/echo/v4"
)

func CreateUser(ctx echo.Context, request viewmodel.SignupRequest) (retVal viewmodel.SignupResponse, err error) {
	user, err := service.CreateUser(ctx, request.ToEntity())
	if err != nil {
		return retVal, err
	}

	return viewmodel.GenerateSignupResponse(user), nil
}
