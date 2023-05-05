package main

import (
	"miniproject_go_wardahfdn/app/infrastructure/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Database Configuration
	db, err := database.NewDatabase()
	if err != nil {
		panic(err)
	}
	// Auto Migrate
	err = db.AutoMigrate()
	if err != nil {
		panic(err)
	}

	// Setup Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	jsonMiddleware := middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		// Do something with request and response body
	})
	e.Use(jsonMiddleware)

	// Start Server
	port := ":8080"
	e.Logger.Fatal(e.Start(port))
}
