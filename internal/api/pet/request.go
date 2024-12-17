package pet

import (
	"github.com/cjmarkham/hexplate/cmd/api"
	"github.com/cjmarkham/hexplate/internal/domain/pet"
)

func RequestToDTO(request api.CreatePetRequestObject) pet.DTO {
	data := request.Body.Data

	return pet.DTO{
		Name: data.Attributes.Name,
	}
}
