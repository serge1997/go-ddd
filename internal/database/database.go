package database

import "gorm.io/gorm"

var Db *gorm.DB

func New() error {
	db, err := gorm.Open(&gorm.Config{})
	if err != nil {
		return err
	}
	Db = db
	return nil
}
