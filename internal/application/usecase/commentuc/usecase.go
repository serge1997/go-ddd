package commentuc

import "github.com/serge1997/go-ddd/internal/domain/comment"

type UseCase struct {
	Create     CreateUseCase
	ListByPost ListByPostUseCase
}

func New(repo comment.Repository) UseCase {
	return UseCase{
		Create:     NewCreateUseCase(repo),
		ListByPost: NewListByPostUseCase(repo),
	}
}
