package user

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type (
	UserModelImpl interface {
		FindByID(id string) User
		FindAll() []User
		Store(name string, email string) User
		Delete(id string) User
	}

	User struct {
		// gorm.Model
		ID        uint       `gorm:"primary_key" json:"id" db:"id"`
		Name      string     `json:"name" db:"name"`
		Email     string     `json:"email" db:"email"`
		CreatedAt time.Time  `json:"created_at" db:"created_at"`
		UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
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

func (u *UserModel) Store(name string, email string) User {
	user := User{Name: name, Email: email}
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
