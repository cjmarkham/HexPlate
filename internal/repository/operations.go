package repository

import (
	"context"
)

type Operations interface {
	InsertOne(context.Context, interface{}) (interface{}, error)
}
