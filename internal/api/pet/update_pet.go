package pet

import (
	"context"
	"github.com/cjmarkham/hexplate/cmd/api"
)

func (h Handlers) UpdatePet(ctx context.Context, request api.UpdatePetRequestObject) (api.UpdatePetResponseObject, error) {
	return api.UpdatePet200JSONResponse{}, nil
}
