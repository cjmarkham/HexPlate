package pet

import (
	"github.com/cjmarkham/hexplate/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
	operations repository.Operations
}

func ProvideRepository(collection *mongo.Collection, operations repository.Operations) Repository {
	return Repository{
		collection: collection,
		operations: operations,
	}
}

func ProvideCollection(db *mongo.Database) *mongo.Collection {
	return db.Collection("pets")
}
