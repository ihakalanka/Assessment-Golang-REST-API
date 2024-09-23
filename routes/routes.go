package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/ihakalanka/Assessment-Golang-REST-API/middleware"
)

func Routes(e *echo.Echo)  {
	v1 := e.Group("/api/v1")

	UserRoute(v1.Group("/user", middleware.JWTMiddleware))
	AuthRoute(v1.Group("/auth"))
}