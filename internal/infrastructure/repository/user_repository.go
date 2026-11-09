package repository

import (
	"context"
	"fmt"

	"github.com/serge1997/go-ddd/internal/domain/user"
	"github.com/serge1997/go-ddd/internal/infrastructure/mapper"
	"github.com/serge1997/go-ddd/internal/infrastructure/model"
	"github.com/serge1997/go-ddd/internal/shared/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) FindAll(ctx context.Context) ([]*user.User, error) {
	var users []model.UserModel
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return mapper.UserModelsToDomainCollection(users), nil
}
func (u *userRepository) Save(ctx context.Context, user *user.User) (*user.User, error) {
	data := mapper.UserToModel(user)
	err := u.db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return nil, err
	}
	return mapper.UserToDomain(data), nil
}
func (u *userRepository) Update(ctx context.Context, user *user.User) (*user.User, error) {
	var model = mapper.UserToModel(user)
	if err := u.db.WithContext(ctx).Save(&model).Error; err != nil {
		return nil, fmt.Errorf("error on update user: %w", err)
	}
	return mapper.UserToDomain(model), nil
}

func (u *userRepository) Find(ctx context.Context, id int) (*user.User, error) {
	var user model.UserModel
	if err := u.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, fmt.Errorf("Post %w: %s", domain.ErrNotFound, err)
	}
	return mapper.UserToDomain(user), nil
}
