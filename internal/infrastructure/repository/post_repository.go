package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/serge1997/go-ddd/internal/domain/post"
	"github.com/serge1997/go-ddd/internal/shared/domain"
)

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) postRepository {
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
