package models

import (
	"github.com/jinzhu/gorm"
	"github.com/khadyCi/bloober/pkg/config"
)

//var db *gorm.DB

type TypeTask struct {
	gorm.Model
	TaskTypeName string `gorm:"" json:"name"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&TypeTask{})

}

func (b *TypeTask) CreateTypeTask() *TypeTask {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllTypeTasks() []TypeTask {
	var TypeTasks []TypeTask
	db.Find(&TypeTasks)
	return TypeTasks
}

func GetTypeTaskById(Id int64) (*TypeTask, *gorm.DB) {
	var GetTypeTask TypeTask
	db := db.Where("ID = ?", Id).Find(&GetTypeTask)
	return &GetTypeTask, db
}

func DeleteTypeTask(ID int64) TypeTask {
	var typeTask TypeTask
	db.Where("ID = ?", ID).Delete(typeTask)
	return typeTask
}
