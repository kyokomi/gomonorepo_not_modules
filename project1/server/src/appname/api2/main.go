package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"appname/api2/fuga"
	"appname/common/piyo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		fmt.Println(piyo.Piyo())
		return c.String(http.StatusOK, "Hello, World! api 2 "+fuga.Fuga())
	})
	e.Logger.Fatal(e.Start(":1324"))
}
