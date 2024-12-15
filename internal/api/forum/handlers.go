package forum

import (
	"context"
	"github.com/cjmarkham/hexplate/cmd/api"
)

func (h Handlers) CreateForum(ctx context.Context, request api.CreateForumRequestObject) (api.CreateForumResponseObject, error) {
	dto := RequestToDTO(request)
	forum, err := h.service.Create(ctx, dto)
	if err != nil {
		return nil, err
	}

	return &api.CreateForum201JSONResponse{
		Id:   forum.ID,
		Name: forum.Name,
	}, nil
}
