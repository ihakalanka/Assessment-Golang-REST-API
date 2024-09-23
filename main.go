package main

import (
	"net/http"
	"os"
	"time"
	"github.com/ihakalanka/Assessment-Golang-REST-API/config"
	"github.com/ihakalanka/Assessment-Golang-REST-API/middleware"
	"github.com/ihakalanka/Assessment-Golang-REST-API/routes"
	"github.com/ihakalanka/Assessment-Golang-REST-API/utils"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
    config.ConnectDB()

    utils.InitLogger()

    e := echo.New()

    e.Use(echomiddleware.RemoveTrailingSlash())

    e.Use(echomiddleware.SecureWithConfig(echomiddleware.SecureConfig{
        XSSProtection:         "1; mode=block",
        ContentTypeNosniff:    "nosniff",
        XFrameOptions:         "DENY",
        HSTSMaxAge:            3600,
        ContentSecurityPolicy: "default-src 'self'",
    }))

    e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
    }))

    e.Use(middleware.Logger())

    e.Use(echomiddleware.Recover())

    routes.Routes(e)

    e.GET("/", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]interface{}{
            "status": 200,
            "message": "Welcome to Golang REST API Server",
            "date":   time.Now().Format(time.RFC3339),
        })
    })

    config.InitEnv()

    PORT := os.Getenv("PORT")
    if PORT == "" {
        PORT = "8080"
        utils.Log.Warn("PORT not set in environment variables, defaulting to 8080")
    }

    utils.Log.Infof("Starting server on port %s", PORT)
    e.Logger.Fatal(e.Start(":" + PORT))
}