package postuc

import (
	"github.com/serge1997/go-ddd/internal/domain/post"
	"github.com/serge1997/go-ddd/internal/domain/user"
)

type UseCase struct {
	CreateUsease  CreateUseCase
	UpdateUseCase UpdateUseCase
	GetUseCase    GetUseCase
	ListByAuthor  ListByAuthorUseCase
}

func New(repo post.Repository, userRepository user.Repository) UseCase {
	return UseCase{
		CreateUsease:  NewCreateUseCase(repo, userRepository),
		UpdateUseCase: NewPostUpdateUseCase(repo, userRepository),
		GetUseCase:    NewGetUseCase(repo),
		ListByAuthor:  NewListByAuthorUseCase(repo, userRepository),
	}
}
