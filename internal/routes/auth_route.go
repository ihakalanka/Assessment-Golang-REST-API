package routes

import (
	"github.com/ihakalanka/Assessment-Golang-REST-API/internal/controllers"
	"github.com/labstack/echo/v4"
)

func AuthRoute(e *echo.Group)  {
	e.POST("/login", controllers.Login)
	e.POST("/refresh-token", controllers.RefreshToken)
}