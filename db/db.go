package db

import (
	"log"
	"os/user"

	_ "github.com/go-sql-driver/mysql" // Mysql driver
	"github.com/jinzhu/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "go-example:go-example@tcp(db:3306)/go-example?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&user.User{})

	return db
}
