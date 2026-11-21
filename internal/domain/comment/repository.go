package comment

import (
	"context"

	"github.com/serge1997/go-ddd/internal/application/dto/commentdto"
)

type Repository interface {
	Save(context.Context, *Comment) (*commentdto.CommentResponse, error)
	FindByPost(ctx context.Context, postId int) ([]commentdto.CommentResponse, error)
}
