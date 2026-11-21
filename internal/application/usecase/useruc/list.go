package useruc

import (
	"context"

	"github.com/serge1997/go-ddd/internal/application/dto/userdto"
	"github.com/serge1997/go-ddd/internal/domain/user"
	"github.com/serge1997/go-ddd/internal/shared/result"
)

type ListUseCase struct {
	repository user.Repository
}

func NewListUseCase(repo user.Repository) ListUseCase {
	return ListUseCase{repo}
}

func (u ListUseCase) Execute(ctx context.Context) result.Result[[]userdto.UserResponseDTO] {
	users, err := u.repository.FindAll(ctx)
	if err != nil {
		return result.Fail[[]userdto.UserResponseDTO](err)
	}
	return result.Ok(user.ToResponseCollection(users))
}
