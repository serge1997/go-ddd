package useruc

import "github.com/serge1997/go-ddd/internal/domain/user"

type UseCase struct {
	CreateUseCase CreateUseCase
	UpdateUseCase UpdateUseCase
	GetUseCase    GetUseCase
	ListUseCase   ListUseCase
}

func New(repo user.Repository) UseCase {
	return UseCase{
		GetUseCase:    NewGetUseCase(repo),
		CreateUseCase: NewCreateUseCase(repo),
		ListUseCase:   NewListUseCase(repo),
		UpdateUseCase: NewUpdateUseCase(repo),
	}
}
