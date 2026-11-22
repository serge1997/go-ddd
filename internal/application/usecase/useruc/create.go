package useruc

import (
	"context"
	"errors"

	"github.com/serge1997/go-ddd/internal/application/dto/userdto"
	"github.com/serge1997/go-ddd/internal/domain/user"
	"github.com/serge1997/go-ddd/internal/shared/domain"
	"github.com/serge1997/go-ddd/internal/shared/result"
)

type CreateUseCase struct {
	repository user.Repository
}

func NewCreateUseCase(repo user.Repository) CreateUseCase {
	return CreateUseCase{repo}
}

func (u CreateUseCase) Execute(ctx context.Context, dto userdto.UserCreateDTO) result.Result[userdto.UserResponseDTO] {
	if err := dto.Validate(); err != nil {
		return result.Fail[userdto.UserResponseDTO](err)
	}
	exists, err := u.repository.FindByEmail(ctx, dto.Email)
	if err != nil && !errors.Is(err, domain.ErrNotFound) {
		return result.Fail[userdto.UserResponseDTO](err)
	}
	if exists != nil {
		return result.Fail[userdto.UserResponseDTO](user.ErrEmailExists)
	}
	us := user.New(dto)
	created, err := u.repository.Save(ctx, us)
	if err != nil {
		return result.Fail[userdto.UserResponseDTO](err)
	}
	return result.Ok(user.ToResponse(created))
}
