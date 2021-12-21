package config

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
)

var (
	dbo *gorm.DB
)

func Connect() {
	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_PROTOCOL"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DBNAME"),
	)

	db, err := gorm.Open("mysql", mysqlCredentials)
	dbo = db

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to the database!")
	}
}

func GetDB() *gorm.DB {
	return dbo
}
