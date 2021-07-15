package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	e.HTTPErrorHandler = NewHttpErrorHandler(NewErrorStatusCodeMaps()).Handler


	// Middleware
	e.Use(middleware.Logger())

	// Routes
	e.GET("/users/:id", getUser)

	// Start server
	e.Start(":1323")
}

func getUser(c echo.Context) error {
	var (
		user *User
		err  error
	)
	id := c.Param("id")

	if user, err = FindUser(id); err != nil { //ErrDocumentNotFound
		return err
	}
	return c.JSON(http.StatusOK, user)
}
