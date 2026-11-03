package postuc

import (
	"context"

	"github.com/serge1997/go-ddd/internal/application/dto/postdto"
	entity "github.com/serge1997/go-ddd/internal/domain/post"
	"github.com/serge1997/go-ddd/internal/shared/result"
)

type PostGetUseCase struct {
	repository entity.Repository
}

func NewPostGetUseCase(repo entity.Repository) PostGetUseCase {
	return PostGetUseCase{repo}
}

func (p *PostGetUseCase) Execute(ctx context.Context, id int) result.Result[postdto.PostResponseDTO] {
	post, err := p.repository.Find(ctx, id)
	if err != nil {
		return result.Fail[postdto.PostResponseDTO](err)
	}
	return result.Ok(entity.ToResponse(post))
}
