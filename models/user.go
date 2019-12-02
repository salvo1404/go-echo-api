package user

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

type (
	UserModelImpl interface {
		FindByID(id string) User
		FindAll() []User
		Store(name string) User
		Delete(id string) User
	}

	User struct {
		gorm.Model
		Name string `json:"name" db:"name"`
	}

	UserModel struct {
		db *gorm.DB
	}
)

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

func (u *UserModel) FindByID(id string) User {
	user := User{}

	result := u.db.First(&user, "id = ?", id)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return user
}

func (u *UserModel) FindAll() []User {
	users := []User{}
	u.db.Find(&users)

	return users
}

func (u *UserModel) Store(name string) User {
	user := User{Name: name}
	result := u.db.Create(&user)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return user
}

func (u *UserModel) Delete(id string) User {
	user := User{}
	u.db.First(&user, "id = ?", id)
	result := u.db.Delete(&user)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return user
}
