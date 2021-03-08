package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	// TODO: routeは別ディレクトリに切り出す
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	log.Println("Server running...")
	if err := e.Start(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
