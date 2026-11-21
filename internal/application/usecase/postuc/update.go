package postuc

import (
	"context"

	"github.com/serge1997/go-ddd/internal/application/dto/postdto"
	"github.com/serge1997/go-ddd/internal/domain/post"
	entity "github.com/serge1997/go-ddd/internal/domain/post"
	"github.com/serge1997/go-ddd/internal/domain/user"
	"github.com/serge1997/go-ddd/internal/shared/result"
)

type PostUpdateUseCase struct {
	repository     entity.Repository
	userRepository user.Repository
}

func NewPostUpdateUseCase(repo post.Repository, userRepository user.Repository) PostUpdateUseCase {
	return PostUpdateUseCase{
		repository:     repo,
		userRepository: userRepository,
	}
}

func (p PostUpdateUseCase) Execute(ctx context.Context, dto postdto.PostUpdateDTO) result.Result[postdto.PostResponseDTO] {
	post, err := p.repository.Find(ctx, dto.Id)
	if err != nil {
		return result.Fail[postdto.PostResponseDTO](err)
	}
	author, err := p.userRepository.Find(ctx, post.Id)
	if err != nil {
		return result.Fail[postdto.PostResponseDTO](err)
	}
	if !post.IsAuthor(author) {
		return result.Fail[postdto.PostResponseDTO](entity.ErrNotAuthor)
	}
	post.Update(dto)
	updated, err := p.repository.Update(ctx, post)
	if err != nil {
		return result.Fail[postdto.PostResponseDTO](err)
	}
	return result.Ok(entity.ToResponse(updated))
}
