package pet

import (
	"context"
	"github.com/cjmarkham/hexplate/cmd/api"
)

func (h Handlers) GetPetById(ctx context.Context, request api.GetPetByIdRequestObject) (api.GetPetByIdResponseObject, error) {
	return api.GetPetById200JSONResponse{}, nil
}
