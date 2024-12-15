package forum

//go:generate mockgen -source=service.go -package=forum -destination=service_mock.go

import "context"

type Service interface {
	Create(ctx context.Context, dto DTO) (*Forum, error)
}

func (s service) Create(ctx context.Context, dto DTO) (*Forum, error) {
	return s.repository.Create(ctx, dto.ToForum())
}
