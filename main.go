package main

import (
	"api/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/onboard", handler.SignupHandler)
	e.Start(":8080")
}
