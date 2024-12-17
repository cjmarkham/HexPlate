package pet

import "github.com/cjmarkham/hexplate/internal/domain/pet"

type Handlers struct {
	service pet.Service
}

func ProvideHandlers(service pet.Service) Handlers {
	return Handlers{
		service: service,
	}
}
