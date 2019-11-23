package db

import (
	"log"

    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql"
)

func DBConnect() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "go-example:go-example@tcp(db:3306)/go-example")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

