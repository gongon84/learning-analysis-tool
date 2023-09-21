package main

import (
	"net/http"
	"strconv"

	"github.com/gongon84/learning-golangci/go/domains"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "1 + 4 = "+strconv.Itoa(domains.Add(1, 4)))
	})
	e.Logger.Fatal(e.Start(":8080"))
}
