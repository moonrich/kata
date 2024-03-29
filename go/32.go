package main

import (
  "net/http"

  "github.com/labstack/echo"
  "github.com/labstack/echo/engine/standard"
  "github.com/labstack/echo/middleware"
)

func main() {

  e := echo.New()

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!\n")
  })

  e.Run(standard.New(":1323"))
}
