package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
}
func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "Worldaaaa")
}