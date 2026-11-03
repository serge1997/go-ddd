package postuc

import (
	"github.com/serge1997/go-ddd/internal/domain/post"
	"github.com/serge1997/go-ddd/internal/domain/user"
)

type UseCase struct {
	CreateUsease  PostCreateUseCase
	UpdateUseCase PostUpdateUseCase
	GetUseCase    PostGetUseCase
	ListByAuthor  PostListByAuthorUseCase
}

func New(repo post.Repository, userRepository user.Repository) UseCase {
	return UseCase{
		CreateUsease:  NewPostCreateUseCase(repo, userRepository),
		UpdateUseCase: NewPostUpdateUseCase(repo, userRepository),
		GetUseCase:    NewPostGetUseCase(repo),
		ListByAuthor:  NewPostListByAuthorUseCase(repo, userRepository),
	}
}
