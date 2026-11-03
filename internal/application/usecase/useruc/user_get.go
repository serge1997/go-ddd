package useruc

import (
	"context"

	"github.com/serge1997/go-ddd/internal/application/dto/userdto"
	"github.com/serge1997/go-ddd/internal/domain/user"
	"github.com/serge1997/go-ddd/internal/shared/result"
)

type UserGetUseCase struct {
	repository user.Repository
}

func NewUserGetUseCase(repository user.Repository) UserGetUseCase {
	return UserGetUseCase{repository}
}

func (u UserGetUseCase) Execute(ctx context.Context, id int) result.Result[userdto.UserResponseDTO] {
	finded, err := u.repository.Find(ctx, id)
	if err != nil {
		return result.Fail[userdto.UserResponseDTO](err)
	}
	return result.Ok(user.ToResponse(finded))
}
