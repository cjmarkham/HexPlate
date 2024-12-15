package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func ProvideDatabase() *mongo.Database {
	dsn := os.Getenv("MONGO_DSN")
	ctx := context.Background()

	mgo, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		panic(err)
	}

	return mgo.Database("hexplate")
}
