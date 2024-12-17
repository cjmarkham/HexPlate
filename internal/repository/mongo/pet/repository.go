package pet

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func ProvideRepository(collection *mongo.Collection) Repository {
	return Repository{
		collection: collection,
	}
}

func ProvideCollection(db *mongo.Database) *mongo.Collection {
	return db.Collection("pets")
}
