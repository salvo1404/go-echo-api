package main

import (
//     "log"

	"net/http"
	"github.com/labstack/echo"

	"app/db"
	"app/handlers"
	"app/models"
)

func main() {
	app := echo.New()

    // Dependencies
    d := db.DBConnect()
    h := userHandler.NewHandler(user.NewUserModel(d))

    // Routes
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Go echo API !")
	})
	app.GET("/users", h.GetIndex)
	app.POST("/users", h.Save)
    app.GET("/users/:id", h.GetDetail)

    // Application start
	app.Logger.Fatal(app.Start(":1323"))
}