package pet

//go:generate mockgen -source=service.go -package=pet -destination=service_mock.go

import "context"

type Service interface {
	Create(ctx context.Context, dto DTO) (*Pet, error)
}
