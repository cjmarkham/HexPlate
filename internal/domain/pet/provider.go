package pet

import "github.com/cjmarkham/hexplate/internal/helpers"

type service struct {
	repository    Repository
	uuidGenerator helpers.DefaultUUIDGenerator
}

func ProvideService(repository Repository, uuidGenerator helpers.DefaultUUIDGenerator) Service {
	return service{
		repository:    repository,
		uuidGenerator: uuidGenerator,
	}
}
