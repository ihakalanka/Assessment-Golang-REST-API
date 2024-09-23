package routes

import (
	"github.com/ihakalanka/Assessment-Golang-REST-API/internal/controllers"
	"github.com/labstack/echo/v4"
	"github.com/ihakalanka/Assessment-Golang-REST-API/internal/middleware"
)

func ProductRoute(g *echo.Group)  {
	g.GET("/product", controllers.GetProducts, middleware.RoleMiddleware("user"))
	g.GET("/product/:id", controllers.GetProduct, middleware.RoleMiddleware("user"))
	g.POST("/product", controllers.CreateProduct, middleware.RoleMiddleware("user"))
	g.PUT("/product/:id", controllers.UpdateProduct, middleware.RoleMiddleware("user"))
	g.DELETE("/product/:id", controllers.DeleteProduct, middleware.RoleMiddleware("user"))
}