package pet

import "github.com/cjmarkham/hexplate/internal/helpers"

type service struct {
	repository    Repository
	uuidGenerator helpers.UUIDGenerator
}

func ProvideService(repository Repository, uuidGenerator helpers.UUIDGenerator) Service {
	return service{
		repository:    repository,
		uuidGenerator: uuidGenerator,
	}
}
