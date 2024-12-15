package forum

import (
	"github.com/cjmarkham/hexplate/internal/domain/forum"
	"github.com/google/uuid"
)

type Document struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

func NewDocument(forum forum.Forum) Document {
	return Document{
		ID:   uuid.New().String(),
		Name: forum.Name,
	}
}
