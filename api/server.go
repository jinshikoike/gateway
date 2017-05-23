package main

import (
	"github.com/jinshikoike/gateway/api/routes"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	routes.Init(e)
	e.Logger.Fatal(e.Start(":8080"))
}
