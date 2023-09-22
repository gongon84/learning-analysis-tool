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
		return c.HTML(http.StatusOK, "<html><body>1 + 4 = "+strconv.Itoa(domains.Add(1, 4))+"</body></html>")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
