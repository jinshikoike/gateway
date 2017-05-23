package routes

import (
	"github.com/jinshikoike/gateway/api/controller"
	"github.com/labstack/echo"
)

// Init routeing
// TODO: support other method
func Init(e *echo.Echo) {
	e.GET("/launch/*", controller.Test)
	e.GET("/*", controller.Caller)
}
