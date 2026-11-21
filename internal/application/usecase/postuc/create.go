package postuc

import (
	"context"

	"github.com/serge1997/go-ddd/internal/application/dto/postdto"
	"github.com/serge1997/go-ddd/internal/domain/post"
	"github.com/serge1997/go-ddd/internal/domain/user"
	"github.com/serge1997/go-ddd/internal/shared/result"
)

type CreateUseCase struct {
	repository     post.Repository
	userRepository user.Repository
}

func NewCreateUseCase(repo post.Repository, userRepository user.Repository) CreateUseCase {
	return CreateUseCase{
		repository:     repo,
		userRepository: userRepository,
	}
}

func (p CreateUseCase) Execute(ctx context.Context, dto postdto.PostCreateDTO) result.Result[postdto.PostResponseDTO] {
	auth, err := p.userRepository.Find(ctx, dto.AuthorId)
	if err != nil {
		return result.Fail[postdto.PostResponseDTO](err)
	}
	newPost := post.New(dto, auth.Id)
	saved, err := p.repository.Save(ctx, newPost)
	if err != nil {
		return result.Fail[postdto.PostResponseDTO](err)
	}
	return result.Ok(post.ToResponse(saved))
}
