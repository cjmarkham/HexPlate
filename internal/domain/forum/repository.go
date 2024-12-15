package forum

//go:generate mockgen -source=repository.go -package=forum -destination=repository_mock.go

import "context"

type Repository interface {
	Create(ctx context.Context, forum Forum) (*Forum, error)
}
