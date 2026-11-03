package post

import (
	"context"

	"github.com/serge1997/go-ddd/internal/shared/domain"
)

type Repository interface {
	domain.Saver[Post]
	domain.Finder[Post]
	domain.Updater[Post]
	FindByAuthor(ctx context.Context, authorId int) ([]*Post, error)
}
