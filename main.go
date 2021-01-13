package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	livereload "github.com/mattn/echo-livereload"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo unasuke ")
	})
	e.GET("/users/:id", getUser)
	e.Logger.Fatal(e.Start(":8080"))

	e.Use(middleware.Static("assets"))
	e.Use(livereload.LiveReload())
	
	//e.POST("/users", saveUser)
	//e.GET("/users/:id", getUser)
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
