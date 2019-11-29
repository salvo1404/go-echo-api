package main

import (
	//     "log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"

	"app/db"
	"app/handlers"
	"app/models"
)

func main() {
	app := echo.New()

	// Config Logger
	app.Use(middleware.Logger())
	// 	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 		Format: "method=${method}, uri=${uri}, status=${status}\n",
	// 	}))

	// Dependencies
	d := db.DBConnect()
	h := userHandler.NewHandler(user.NewUserModel(d))

	// BasicAuth
	//     admin := e.Group("/admin", middleware.BasicAuth())
	// For invalid credentials
	//     		return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

	// Routes
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Go echo API !")
	})
	app.GET("/users", h.GetIndex)
	app.POST("/users", h.Save)
	app.GET("/users/:id", h.GetDetail)
	app.DELETE("/users/:id", h.Delete)

	// Application start
	app.Logger.Fatal(app.Start(":1323"))
}
