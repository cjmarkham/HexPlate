package pet

import (
	"context"
	"github.com/cjmarkham/hexplate/internal/domain/pet"
	"github.com/google/uuid"
)

func (r Repository) Create(ctx context.Context, p pet.Pet) (*pet.Pet, error) {
	doc := NewDocument(p)
	_, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}
	p.ID = uuid.MustParse(doc.ID)
	return &p, nil
}
