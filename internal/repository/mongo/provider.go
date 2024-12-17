package mongo

import (
	"context"
	"github.com/cjmarkham/hexplate/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func ProvideDatabase() *mongo.Database {
	dsn := os.Getenv("MONGO_DSN")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mgo, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		panic(err)
	}

	return mgo.Database("hexplate")
}

type Operations struct {
	collection *mongo.Collection
}

func ProvideOperations(collection *mongo.Collection) Operations {
	return Operations{
		collection: collection,
	}
}

var _ repository.Operations = (*Operations)(nil)
