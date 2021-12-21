package models

import (
	"github.com/jinzhu/gorm"
	"github.com/khadyCi/bloober/pkg/config"
)

//var db *gorm.DB
//Allow look like permission

type Allow struct {
	gorm.Model
	Name        string ` json:"name"`
	Description string ` json:"description"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Allow{})
}

func (b *Allow) CreateAllow() *Allow {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllAllows() []Allow {
	var Allows []Allow
	db.Find(&Allows)
	return Allows
}

func GetAllowById(Id int64) (*Allow, *gorm.DB) {
	var GetAllow Allow
	db := db.Where("ID = ?", Id).Find(&GetAllow)
	return &GetAllow, db
}

func DeleteAllow(ID int64) Allow {
	var allow Allow
	db.Where("ID = ?", ID).Delete(allow)
	return allow
}
