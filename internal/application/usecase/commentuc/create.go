package commentuc

import (
	"context"

	"github.com/serge1997/go-ddd/internal/application/dto/commentdto"
	"github.com/serge1997/go-ddd/internal/domain/comment"
	"github.com/serge1997/go-ddd/internal/shared/result"
)

type CreateUseCase struct {
	repo comment.Repository
}

func NewCreateUseCase(repo comment.Repository) CreateUseCase {
	return CreateUseCase{repo}
}

func (c *CreateUseCase) Execute(ctx context.Context) result.Result[commentdto.Response] {
	resp, err := c.repo.Save(ctx, &comment.Comment{})
	if err != nil {
		return result.Fail[commentdto.Response](err)
	}
	return result.Ok(comment.ToResponseFromRaw(*resp))
}
