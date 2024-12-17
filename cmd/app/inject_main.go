//go:build wireinject
// +build wireinject

package main

import (
	forumApi "github.com/cjmarkham/hexplate/internal/api/pet"
	forumService "github.com/cjmarkham/hexplate/internal/domain/pet"
	"github.com/cjmarkham/hexplate/internal/repository/mongo"
	forumMongo "github.com/cjmarkham/hexplate/internal/repository/mongo/pet"
	"github.com/google/wire"
)

var apiHandlers = wire.NewSet(
	forumApi.ProvideHandlers,
)

var domainHandlers = wire.NewSet(
	forumService.ProvideService,
)

var repositoryHandlers = wire.NewSet(
	mongo.ProvideDatabase,
	forumMongo.ProvideRepository,
	forumMongo.ProvideCollection,
	wire.Bind(new(forumService.Repository), new(forumMongo.Repository)),
)

func injectApp() forumApi.Handlers {
	panic(
		wire.Build(
			apiHandlers,
			domainHandlers,
			repositoryHandlers,
		),
	)
}
