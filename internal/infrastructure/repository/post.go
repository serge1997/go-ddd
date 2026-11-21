package repository

import (
	"context"
	"fmt"

	"github.com/serge1997/go-ddd/internal/application/dto/postdto"
	"github.com/serge1997/go-ddd/internal/domain/post"
	"github.com/serge1997/go-ddd/internal/infrastructure/mapper"
	"github.com/serge1997/go-ddd/internal/infrastructure/model"
	"github.com/serge1997/go-ddd/internal/shared/domain"
	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) postRepository {
	return postRepository{db}
}

func (u postRepository) FindAll(ctx context.Context) ([]*postdto.Row, error) {
	var rows []*postdto.Row
	err := u.db.WithContext(ctx).Table("posts as p").
		Select("p.*, u.Name as author_name").
		Joins("JOIN users as u on u.id = p.author_id").
		Order("id desc").
		Scan(rows).
		Error
	if err != nil {
		return nil, err
	}
	return rows, nil
}
func (u postRepository) Save(ctx context.Context, post *post.Post) (*postdto.Row, error) {
	m := model.NewPostModel(post)
	err := u.db.WithContext(ctx).Save(m).Preload("Author").Error
	if err != nil {
		return nil, err
	}
	return mapper.ToPostRaw(m), nil
}

func (u postRepository) Update(ctx context.Context, post *post.Post) (*postdto.Row, error) {
	m, err := model.UpdatePost(post)
	if err != nil {
		return nil, err
	}
	err = u.db.WithContext(ctx).Save(m).Error
	if err != nil {
		return nil, err
	}
	return mapper.ToPostRaw(m), nil
}

func (u postRepository) Find(ctx context.Context, id int) (*postdto.Row, error) {
	var p model.PostModel
	err := u.db.WithContext(ctx).Preload("Author").First(&p, id).Error
	if err != nil {
		return nil, fmt.Errorf("Post %w", domain.ErrNotFound)
	}
	return mapper.ToPostRaw(&p), nil
}

func (u postRepository) FindByAuthor(ctx context.Context, authorId int) ([]postdto.Row, error) {
	var rows []postdto.Row
	err := u.db.WithContext(ctx).Table("posts as p").
		Select("p.*, u.Name as author_name").
		Joins("JOIN users as u on u.id = p.author_id").
		Where("p.author_id = ?", authorId).
		Order("id desc").
		Scan(&rows).
		Error
	if err != nil {
		return nil, err
	}
	return rows, nil
}
