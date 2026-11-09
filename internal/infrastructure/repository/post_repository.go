package repository

import (
	"context"
	"fmt"

	"github.com/serge1997/go-ddd/internal/domain/post"
	"github.com/serge1997/go-ddd/internal/shared/domain"
	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) postRepository {
	return postRepository{db}
}

func (u postRepository) FindAll(ctx context.Context) ([]*post.Post, error) {
	return []*post.Post{}, nil
}
func (u postRepository) Save(ctx context.Context, post *post.Post) (*post.Post, error) {
	return post, nil
}

func (u postRepository) Update(ctx context.Context, post *post.Post) (*post.Post, error) {
	return post, nil
}

func (u postRepository) Find(ctx context.Context, id int) (*post.Post, error) {
	return nil, fmt.Errorf("Post %w", domain.ErrNotFound)
}

func (u postRepository) FindByAuthor(ctx context.Context, authorId int) ([]*post.Post, error) {
	return []*post.Post{}, nil
}
