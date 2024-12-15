package forum

import (
	"context"
	"github.com/cjmarkham/hexplate/internal/domain/forum"
	"github.com/google/uuid"
)

func (r Repository) Create(ctx context.Context, forum forum.Forum) (*forum.Forum, error) {
	doc := NewDocument(forum)
	_, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}
	forum.ID = uuid.MustParse(doc.ID)
	return &forum, nil
}
