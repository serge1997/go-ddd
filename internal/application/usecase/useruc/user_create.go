package useruc

import (
	"context"

	"github.com/serge1997/go-ddd/internal/application/dto/userdto"
	"github.com/serge1997/go-ddd/internal/domain/user"
	"github.com/serge1997/go-ddd/internal/shared/result"
)

type UserCreateUseCase struct {
	Repository user.Repository
}

func NewCreateUseCase(repo user.Repository) UserCreateUseCase {
	return UserCreateUseCase{repo}
}

func (u UserCreateUseCase) Execute(ctx context.Context, dto userdto.UserCreateDTO) result.Result[userdto.UserResponseDTO] {
	if err := dto.Validate(); err != nil {
		return result.Fail[userdto.UserResponseDTO](err)
	}
	us := user.New(dto)
	created, err := u.Repository.Save(ctx, us)
	if err != nil {
		return result.Fail[userdto.UserResponseDTO](err)
	}
	return result.Ok(user.ToResponse(created))
}
