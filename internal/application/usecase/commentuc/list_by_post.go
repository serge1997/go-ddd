package commentuc

import (
	"context"

	"github.com/serge1997/go-ddd/internal/application/dto/commentdto"
	"github.com/serge1997/go-ddd/internal/domain/comment"
	"github.com/serge1997/go-ddd/internal/shared/result"
)

type ListByPostUseCase struct {
	repo comment.Repository
}

func NewListByPostUseCase(repo comment.Repository) ListByPostUseCase {
	return ListByPostUseCase{repo}
}

func (c *ListByPostUseCase) Execute(ctx context.Context) result.Result[[]commentdto.Response] {
	rows, err := c.repo.FindByPost(ctx, 34)
	if err != nil {
		return result.Fail[[]commentdto.Response](err)
	}
	return result.Ok(comment.ToResponseList(rows))
}
