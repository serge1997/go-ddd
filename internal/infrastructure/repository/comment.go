package repository

import (
	"github.com/serge1997/go-ddd/internal/application/dto/commentdto"
	"github.com/serge1997/go-ddd/internal/domain/comment"
	"github.com/serge1997/go-ddd/internal/infrastructure/mapper"
	"github.com/serge1997/go-ddd/internal/infrastructure/model"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) commentRepository {
	return commentRepository{db}
}

func (c commentRepository) Save(ctx context.Context, com *comment.Comment) (*commentdto.Raw, error) {
	m := model.NewCommentModel(com)
	err := c.db.WithContext(ctx).Preload("Author").Save(m).Error
	if err != nil {
		return nil, err
	}
	return mapper.ToCommentRaw(m), nil
}

func (c commentRepository) FindByPost(ctx context.Context, postId int) ([]commentdto.Raw, error) {
	var rows []commentdto.Raw
	err := c.db.WithContext(ctx).Table("comments as c").
		Select(`c.id, c.author_id, u.name as author_name,
			c.content, c.created_at, c.updated_at
		`).
		Joins("users as u on u.id = c.author_id").
		Where("post_id = ?", postId).
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}
	return rows, nil
}
