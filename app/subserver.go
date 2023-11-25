package main

import {
	"fmt"
}

func Hello(e){

	e.GET("/hello", func(c echo.Context) error{
		return c.Render(http.StatusOK, "hello", "Worldaaaa")
	})
}