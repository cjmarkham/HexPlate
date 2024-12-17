package pet

import (
	"context"
	"github.com/cjmarkham/hexplate/cmd/api"
)

func (h Handlers) CreatePet(ctx context.Context, request api.CreatePetRequestObject) (api.CreatePetResponseObject, error) {
	dto := RequestToDTO(request)
	pet, err := h.service.Create(ctx, dto)
	if err != nil {
		return nil, err
	}

	return &api.CreatePet201JSONResponse{
		Data: api.PetResponse{
			Id:   pet.ID,
			Type: api.PetResponseTypePet,
			Attributes: api.PetResponseAttributes{
				Name: pet.Name,
			},
		},
	}, nil
}
