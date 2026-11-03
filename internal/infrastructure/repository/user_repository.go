package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/serge1997/go-ddd/internal/domain/user"
	"github.com/serge1997/go-ddd/internal/shared/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) FindAll(ctx context.Context) ([]*user.User, error) {
	return []*user.User{}, nil
}
func (u *userRepository) Save(ctx context.Context, user *user.User) (*user.User, error) {
	return user, nil
}
func (u *userRepository) Update(ctx context.Context, user *user.User) (*user.User, error) {
	return user, nil
}

func (u *userRepository) Find(ctx context.Context, id int) (*user.User, error) {
	return nil, fmt.Errorf("Post %w", domain.ErrNotFound)
}
