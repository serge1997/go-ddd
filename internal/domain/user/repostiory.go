package user

import (
	"context"

	"github.com/serge1997/go-ddd/internal/shared/domain"
)

type Repository interface {
	domain.Saver[User]
	domain.Updater[User]
	domain.Finder[User]
	domain.Lister[User]
	FindByEmail(context.Context, string) (*User, error)
}
