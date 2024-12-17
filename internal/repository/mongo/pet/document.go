package pet

import (
	"github.com/cjmarkham/hexplate/internal/domain/pet"
	"github.com/google/uuid"
)

type Document struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

func NewDocument(p pet.Pet) Document {
	return Document{
		ID:   uuid.New().String(),
		Name: p.Name,
	}
}
