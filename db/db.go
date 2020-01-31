package db

import (
	user "github.com/salvo1404/go-echo-api/models"
	"log"

	"github.com/go-sql-driver/mysql" // Mysql driver
	"github.com/jinzhu/gorm"
)

func Connect() *gorm.DB {
	DBMS := "mysql"
	mySQLConfig := &mysql.Config{
		User:                 "go-example",
		Passwd:               "go-example",
		Net:                  "tcp",
		Addr:                 "db:3306",
		DBName:               "go-example",
		AllowNativePasswords: true,
		Params: map[string]string{
			"parseTime": "true",
		},
	}

	db, err := gorm.Open(DBMS, mySQLConfig.FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&user.User{})

	return db
}
