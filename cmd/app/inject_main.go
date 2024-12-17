//go:build wireinject
// +build wireinject

package main

import (
	petApi "github.com/cjmarkham/hexplate/internal/api/pet"
	petService "github.com/cjmarkham/hexplate/internal/domain/pet"
	"github.com/cjmarkham/hexplate/internal/helpers"
	"github.com/cjmarkham/hexplate/internal/repository"
	"github.com/cjmarkham/hexplate/internal/repository/mongo"
	"github.com/cjmarkham/hexplate/internal/repository/mongo/pet"
	"github.com/google/wire"
)

var apiHandlers = wire.NewSet(
	petApi.ProvideHandlers,
)

var domainHandlers = wire.NewSet(
	helpers.ProvideUUIDGenerator,
	petService.ProvideService,
)

var repositoryHandlers = wire.NewSet(
	mongo.ProvideDatabase,
	pet.ProvideCollection,
	mongo.ProvideOperations,
	wire.Bind(new(repository.Operations), new(mongo.Operations)),
	pet.ProvideRepository,
	wire.Bind(new(petService.Repository), new(pet.Repository)),
)

func injectApp() petApi.Handlers {
	panic(
		wire.Build(
			apiHandlers,
			domainHandlers,
			repositoryHandlers,
		),
	)
}
