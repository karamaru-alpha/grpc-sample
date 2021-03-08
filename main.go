package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/karamaru-alpha/grpc-sample/config"
)

func main() {
	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	log.Println("Server running...")
	if err := e.Start(":" + config.Port()); err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
