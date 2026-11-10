package model

import "gorm.io/gorm"

type PostModel struct {
	gorm.Model
	Id       int       `gorm:"column:id,primaryKey"`
	Title    string    `gorm:"column:title"`
	Content  string    `gorm:"column:title"`
	AuthorId uint      `gorm:"column:author_id"`
	Author   UserModel `gorm:"foreignKey:AuthorId"`
}

func (PostModel) TableName() string {
	return "posts"
}
