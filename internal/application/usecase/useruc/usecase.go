package useruc

import "github.com/serge1997/go-ddd/internal/domain/user"

type UseCase struct {
	CreateUseCase UserCreateUseCase
	UpdateUseCase UserUpdateUseCase
	GetUseCase    UserGetUseCase
	ListUseCase   UserListUseCase
}

func New(repo user.Repository) UseCase {
	return UseCase{
		GetUseCase:    NewUserGetUseCase(repo),
		CreateUseCase: NewCreateUseCase(repo),
		ListUseCase:   NewUserListUseCase(repo),
		UpdateUseCase: NewUserUpdateUseCase(repo),
	}
}
