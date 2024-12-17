package repository

//go:generate mockgen -source=operations.go -package=repository -destination=operations_mock.go

import (
	"context"
)

type Operations interface {
	InsertOne(context.Context, interface{}) (interface{}, error)
}
