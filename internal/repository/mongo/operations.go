package mongo

import (
	"context"
)

func (o Operations) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	return o.collection.InsertOne(ctx, document)
}
