package pet

//go:generate mockgen -source=repository.go -package=pet -destination=repository_mock.go

import "context"

type Repository interface {
	Create(context.Context, Pet) (*Pet, error)
}
