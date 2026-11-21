package comment

import (
	"context"

	"github.com/serge1997/go-ddd/internal/application/dto/commentdto"
)

type Repository interface {
	Save(context.Context, *Comment) (*commentdto.Raw, error)
	FindByPost(ctx context.Context, postId int) ([]commentdto.Raw, error)
}
