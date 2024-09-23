package routes

import (
	"github.com/ihakalanka/Assessment-Golang-REST-API/internal/controllers"
	"github.com/labstack/echo/v4"
	"github.com/ihakalanka/Assessment-Golang-REST-API/internal/middleware"
)

func UserRoute(e *echo.Group)  {
	e.POST("/", controllers.RegisterUser, middleware.RoleMiddleware("admin"))
	e.GET("/", controllers.GetAllUsers)
	e.DELETE("/:id", controllers.DeleteUser, middleware.RoleMiddleware("admin"))
	e.PUT("/:id", controllers.UpdateUser, middleware.RoleMiddleware("admin"))
}