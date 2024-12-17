// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/cjmarkham/hexplate/internal/api/pet"
	pet3 "github.com/cjmarkham/hexplate/internal/domain/pet"
	"github.com/cjmarkham/hexplate/internal/helpers"
	"github.com/cjmarkham/hexplate/internal/repository"
	"github.com/cjmarkham/hexplate/internal/repository/mongo"
	pet2 "github.com/cjmarkham/hexplate/internal/repository/mongo/pet"
	"github.com/google/wire"
)

// Injectors from inject_main.go:

func injectApp() pet.Handlers {
	database := mongo.ProvideDatabase()
	collection := pet2.ProvideCollection(database)
	operations := mongo.ProvideOperations(collection)
	repository := pet2.ProvideRepository(collection, operations)
	defaultUUIDGenerator := helpers.ProvideUUIDGenerator()
	service := pet3.ProvideService(repository, defaultUUIDGenerator)
	handlers := pet.ProvideHandlers(service)
	return handlers
}

// inject_main.go:

var apiHandlers = wire.NewSet(pet.ProvideHandlers)

var domainHandlers = wire.NewSet(helpers.ProvideUUIDGenerator, pet3.ProvideService)

var repositoryHandlers = wire.NewSet(mongo.ProvideDatabase, pet2.ProvideCollection, mongo.ProvideOperations, wire.Bind(new(repository.Operations), new(mongo.Operations)), pet2.ProvideRepository, wire.Bind(new(pet3.Repository), new(pet2.Repository)))
