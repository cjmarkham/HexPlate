package pet

import (
	"context"
	"github.com/cjmarkham/hexplate/cmd/api"
)

func (h Handlers) DeletePet(ctx context.Context, request api.DeletePetRequestObject) (api.DeletePetResponseObject, error) {
	return api.DeletePet204Response{}, nil
}
