package main

import (
	"html/template"
  "net/http"
  "github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
}
func Hello(c *CustomContext) {
	return c.Render(http.StatusOK, "hello", "Worldaaaa")
}