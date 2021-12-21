package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/khadyCi/bloober/pkg/config"
)

//var db *gorm.DB
//Allow look like permission

type UserTask struct {
	UserID    int    `gorm:"primaryKey"`
	TaskID    int    `gorm:"primaryKey"`
	FileName  string ``
	User      User   `json:"user"`
	Task      Task   `json:"task"`
	CreatedAt time.Time
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&UserAllow{})
}

func GetUserTasks(UserID int64) ([]UserTask, *gorm.DB) {
	var UserTask []UserTask
	db := db.Where("UserID = ?", UserID).Find(&UserTask)
	return UserTask, db
}

func (b *User) CreateUserTask() *User {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
