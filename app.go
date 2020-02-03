package main

import (
	// "log"
	// "fmt"

	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/friendsofgo/graphiql"
	"github.com/salvo1404/go-echo-api/db"
	echoApiGraphql "github.com/salvo1404/go-echo-api/graphql"
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

	// Database
	d := db.Connect()
	d.LogMode(true)
	defer d.Close()

	h := userHandler.NewHandler(user.NewUserModel(d))

	// BasicAuth
	//     admin := e.Group("/admin", middleware.BasicAuth())
	// For invalid credentials
	//     		return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

	// Routes
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Go echo API !")
	})

	// Graphql endpoint here
	graphHandler, err := echoApiGraphql.NewGraphHandler(d)
	if err != nil {
		log.Fatalln(err)
	}
	app.POST("/graphql", echo.WrapHandler(graphHandler))

	// Endpoint GraphQl interface http://localhost:1323/graphiql
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		log.Fatalln(err)
	}
	app.GET("/graphiql", echo.WrapHandler(graphiqlHandler))

	app.GET("/users", h.GetIndex)
	app.POST("/users", h.Save)
	app.GET("/users/:id", h.GetDetail)
	app.DELETE("/users/:id", h.Delete)

	// Application start
	app.Logger.Fatal(app.Start(":1323"))
}
