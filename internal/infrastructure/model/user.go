package model

import (
	"github.com/serge1997/go-ddd/internal/domain/user"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Id    int    `gorm:"column:id,primaryKey"`
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
}

func NewUserModel(u *user.User) UserModel {
	return UserModel{
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email.Value(),
	}
}
