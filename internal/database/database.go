package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

// mysql://root:@localhost:3306/goddd
func init() {
	dsn := "root:@tcp(localhost:3306)/goddd?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("Database connected")
	Db = db
}
