package useruc

import (
	"context"

	"github.com/serge1997/go-ddd/internal/application/dto/userdto"
	entity "github.com/serge1997/go-ddd/internal/domain/user"
	"github.com/serge1997/go-ddd/internal/shared/result"
)

type UserUpdateUseCase struct {
	repository entity.Repository
}

func NewUserUpdateUseCase(repo entity.Repository) UserUpdateUseCase {
	return UserUpdateUseCase{repo}
}

func (u UserUpdateUseCase) Execute(ctx context.Context, dto userdto.UserUpdateDTO) result.Result[userdto.UserResponseDTO] {
	user, err := u.repository.Find(ctx, dto.Id)
	if err != nil {
		return result.Fail[userdto.UserResponseDTO](err)
	}
	user.Update(dto)
	updated, err := u.repository.Update(ctx, user)
	if err != nil {
		return result.Fail[userdto.UserResponseDTO](err)
	}
	return result.Ok(entity.ToResponse(updated))
}
