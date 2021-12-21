package models

import (
	"github.com/jinzhu/gorm"
	"github.com/khadyCi/bloober/pkg/config"
)

//var db *gorm.DB

type Task struct {
	gorm.Model
	Title        string `gorm:"" json:"title"`
	TaskTypeName string `gorm:"type_varchar(50); not null" json:"taskTypeName"`
	Importance   string `gorm:"type_varchar(10); not null" json:"importance"`
	Description  string `gorm:"type_varchar(500); not null" json:"description"`
	PubDate      string `gorm:"type_varchar(20); not null" json:"PubDate"`
	FinalDate    string `gorm:"type_varchar(20); not null" json:"FinalDate"`
	OwnerID      int    `json:"owner_id"`
	Users        []User `gorm:"many2many:user_tasks" json:"users"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Task{})
}

func (b *Task) CreateTask() *Task {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllTasks() []Task {
	var Tasks []Task
	db.Preload("Users").Find(&Tasks)
	return Tasks
}

func GetTaskById(Id int64) (*Task, *gorm.DB) {
	var GetTask Task
	db := db.Where("ID = ?", Id).Find(&GetTask)
	return &GetTask, db
}

func GetTaskByOwnerId(owner_id int64) (*Task, *gorm.DB) {
	var GetTask Task
	db := db.Where("OwnerID = ?", owner_id).Find(&GetTask)
	return &GetTask, db
}

func DeleteTask(ID int64) Task {
	var task Task
	db.Where("ID = ?", ID).Delete(task)
	return task
}
