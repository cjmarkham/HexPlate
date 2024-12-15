package forum

import "github.com/cjmarkham/hexplate/internal/domain/forum"

type Handlers struct {
	service forum.Service
}

func ProvideHandlers(service forum.Service) Handlers {
	return Handlers{
		service: service,
	}
}
