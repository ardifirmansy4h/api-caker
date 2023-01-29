package main

import (
	"caker/app/config"

	"caker/app/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Connect()

	e := echo.New()
	routes.SetRoute(e)
	e.Logger.Fatal(e.Start(":2325"))

}
