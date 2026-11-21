package post

import (
	"context"

	"github.com/serge1997/go-ddd/internal/application/dto/postdto"
)

type Repository interface {
	Save(context.Context, *Post) (*postdto.Row, error)
	FindAll(context.Context) ([]*postdto.Row, error)
	Find(context.Context, int) (*postdto.Row, error)
	Update(context.Context, *Post) (*postdto.Row, error)
	FindByAuthor(ctx context.Context, authorId int) ([]postdto.Row, error)
}
