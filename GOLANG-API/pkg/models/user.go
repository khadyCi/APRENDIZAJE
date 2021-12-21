package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/khadyCi/bloober/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB

type User struct {
	gorm.Model

	Dni          string  ` json:"dni"`
	Name         string  ` json:"name"`
	Username     string  `gorm:"not null; uniqueIndex:uid_username" json:"username"`
	Last_name    string  `gorm:"type_varchar(40); not null" json:"last_name"`
	Phone_number string  `gorm:"type_varchar(15); not null" json:"phone_number"`
	Role         string  `gorm:"type_varchar(40); not null" json:"role"`
	Direction    string  `gorm:"type_varchar(90); not null" json:"direction"`
	Section      string  `gorm:"type_varchar(40); not null" json:"section"`
	Email        string  `gorm:"type_varchar(60); not null; unique_index" json:"email"`
	Imagen       string  `gorm:"type_varchar(255); not null" json:"imagen"`
	Postal_code  string  `gorm:"type_varchar(10); not null" json:"postal_code"`
	Password     string  `json:""`
	PasswordHash string  `gorm:"type_varchar(60); not null" json:"passwordHash"`
	Allows       []Allow `gorm:"many2many:user_allows" json:"allows"`
	Tasks        []Task  `gorm:"many2many:user_tasks" json:"tasks"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (b *User) CreateUser() *User {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users) //(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var GetUser User
	db := db.Where("ID = ?", Id).Find(&GetUser)
	return &GetUser, db
}

func DeleteUser(ID int64) User {
	var user User
	db.Where("ID = ?", ID).Delete(user)
	return user
}

func (u *User) GetByUserName(Username string) User {
	var user User
	db.Where("Username = ?", Username).First(&user)
	return user
}

// HashPassword generates a hash of the password and places the result in PasswordHash.
func (u *User) HashPassword() error {
	fmt.Println(u.Password)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PasswordHash = string(passwordHash)

	return nil
}

// PasswordMatch compares HashPassword with the password and returns true if they match.
func (u User) PasswordMatch(password string) bool {
	fmt.Println(u.PasswordHash)
	fmt.Println(password)
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		fmt.Println(err)
	}

	return err == nil
}

//User_router.go
