package domain

import "context"

type Saver[T interface{}] interface {
	Save(ctx context.Context, entity *T) (*T, error)
}
type Finder[T interface{}] interface {
	Find(ctx context.Context, id int) (*T, error)
}
type Lister[T interface{}] interface {
	FindAll(ctx context.Context) ([]*T, error)
}
type Updater[T interface{}] interface {
	Update(ctx context.Context, entity *T) (*T, error)
}
type Deleter[T interface{}] interface {
	Delete(ctx context.Context, id int) error
}
type Repository[T interface{}] interface {
	FindAll(ctx context.Context) ([]*T, error)
	Save(ctx context.Context, entity *T) (*T, error)
}
