package models

import (
	"github.com/jinzhu/gorm"
	"github.com/khadyCi/bloober/pkg/config"
)

//var db *gorm.DB
//Allow look like permission

type UserAllow struct {
	gorm.Model
	UserId  string  `gorm:"primaryKey"`
	AllowId string  `gorm:"primaryKey"`
	User    User    ` json:"user"`
	Allows  []Allow ` json:"allows"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&UserAllow{})
}

func (b *UserAllow) CreateUserAllow() *UserAllow {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllUserAllows() []UserAllow {
	var UserAllows []UserAllow
	db.Preload("User").Preload("Allow").Find(&UserAllows)
	return UserAllows
}

func GetUserAllowById(Id int64) (*UserAllow, *gorm.DB) {
	var GetUserAllow UserAllow
	db := db.Where("ID = ?", Id).Find(&GetUserAllow)
	return &GetUserAllow, db
}

func DeleteUserAllow(ID int64) UserAllow {
	var UserAllow UserAllow
	db.Where("ID = ?", ID).Delete(UserAllow)
	return UserAllow
}
