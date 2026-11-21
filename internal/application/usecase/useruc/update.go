package useruc

import (
	"context"

	"github.com/serge1997/go-ddd/internal/application/dto/userdto"
	entity "github.com/serge1997/go-ddd/internal/domain/user"
	"github.com/serge1997/go-ddd/internal/shared/result"
)

type UpdateUseCase struct {
	repository entity.Repository
}

func NewUpdateUseCase(repo entity.Repository) UpdateUseCase {
	return UpdateUseCase{repo}
}

func (u UpdateUseCase) Execute(ctx context.Context, dto userdto.UserUpdateDTO) result.Result[userdto.UserResponseDTO] {
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
