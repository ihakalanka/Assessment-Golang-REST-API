package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ihakalanka/Assessment-Golang-REST-API/utils"
)

func Logger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339} remote_ip=${remote_ip} method=${method} uri=${uri} status=${status} latency=${latency_human}\n",
		Output: utils.Log.Out,
	})
}