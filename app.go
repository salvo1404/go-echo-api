package main

import (
	// "log"

	// "fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/salvo1404/go-echo-api/db"
	userHandler "github.com/salvo1404/go-echo-api/handlers"
	user "github.com/salvo1404/go-echo-api/models"
)

func main() {
	app := echo.New()

	// fmt.Println("ciao")

	// Config Logger
	app.Use(middleware.Logger())
	// 	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 		Format: "method=${method}, uri=${uri}, status=${status}\n",
	// 	}))

	// Dependencies
	d := db.Connect()
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
