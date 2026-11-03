package postuc

import (
	"context"

	"github.com/serge1997/go-ddd/internal/application/dto/postdto"
	"github.com/serge1997/go-ddd/internal/domain/post"
	"github.com/serge1997/go-ddd/internal/domain/user"
	"github.com/serge1997/go-ddd/internal/shared/result"
)

type PostListByAuthorUseCase struct {
	reposistory     post.Repository
	userReposistory user.Repository
}

func NewPostListByAuthorUseCase(repo post.Repository, userRepository user.Repository) PostListByAuthorUseCase {
	return PostListByAuthorUseCase{
		reposistory:     repo,
		userReposistory: userRepository,
	}
}
func (p PostListByAuthorUseCase) Execute(ctx context.Context, authorId int) result.Result[[]postdto.PostResponseDTO] {
	auth, err := p.userReposistory.Find(ctx, authorId)
	if err != nil {
		return result.Fail[[]postdto.PostResponseDTO](err)
	}
	posts, err := p.reposistory.FindByAuthor(ctx, auth.Id)
	if err != nil {
		return result.Fail[[]postdto.PostResponseDTO](err)
	}
	return result.Ok(post.ToResponseCollection(posts))
}
