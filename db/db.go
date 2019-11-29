package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"app/models"
)

func DBConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "go-example:go-example@tcp(db:3306)/go-example")
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&user.User{})

	return db
}
