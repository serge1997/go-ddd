package model

import (
	"fmt"

	"github.com/serge1997/go-ddd/internal/domain/post"
	"gorm.io/gorm"
)

type PostModel struct {
	gorm.Model
	Id       int       `gorm:"column:id,primaryKey"`
	Title    string    `gorm:"column:title"`
	Content  string    `gorm:"column:title"`
	AuthorId int       `gorm:"column:author_id"`
	Author   UserModel `gorm:"foreignKey:AuthorId"`
}

func (PostModel) TableName() string {
	return "posts"
}

func NewPostModel(p *post.Post) *PostModel {
	return &PostModel{
		Id:       p.Id,
		AuthorId: p.AuthorId,
		Content:  p.Content,
		Title:    p.Title,
	}
}

func UpdatePost(p *post.Post) (*PostModel, error) {
	if p.Id == 0 {
		return nil, fmt.Errorf("Id must be provided for update record")
	}
	return &PostModel{
		Id:       p.Id,
		AuthorId: p.AuthorId,
		Content:  p.Content,
		Title:    p.Title,
	}, nil
}
