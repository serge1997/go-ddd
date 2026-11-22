package useruc

import (
	"context"
	"errors"

	"github.com/serge1997/go-ddd/internal/application/dto/userdto"
	"github.com/serge1997/go-ddd/internal/domain/user"
	"github.com/serge1997/go-ddd/internal/shared/domain"
	"github.com/serge1997/go-ddd/internal/shared/result"
)

type GetUseCase struct {
	repository user.Repository
}

func NewGetUseCase(repository user.Repository) GetUseCase {
	return GetUseCase{repository}
}

func (u GetUseCase) Execute(ctx context.Context, id int) result.Result[userdto.UserResponseDTO] {
	finded, err := u.repository.Find(ctx, id)
	if err != nil && !errors.Is(err, domain.ErrNotFound) {
		return result.Fail[userdto.UserResponseDTO](err)
	}
	if finded == nil {
		return result.Fail[userdto.UserResponseDTO](domain.ErrNotFound)
	}
	return result.Ok(user.ToResponse(finded))
}
