package model

import (
	"time"

	"github.com/serge1997/go-ddd/internal/domain/comment"
)

type CommentModel struct {
	Id        int `gorm:"primaryKey"`
	Content   string
	AuthorId  int
	PostId    int
	CreatedAt time.Time
	Author    UserModel `gorm:"foreignKey:AuthorId"`
}

func (c CommentModel) TableName() string {
	return "comments"
}

func NewCommentModel(c *comment.Comment) *CommentModel {
	return &CommentModel{
		Id:        c.Id,
		AuthorId:  c.AuthorId,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
	}
}
