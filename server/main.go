package main

import (
	"github.com/labstack/echo/v4"
	"server/endpoints/account"
)

func main() {
	e := echo.New()
	account.RegisterHandlers(e)

	e.Logger.Fatal(e.Start(":8080"))
}
