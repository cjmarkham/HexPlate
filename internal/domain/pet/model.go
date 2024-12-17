package pet

import "github.com/google/uuid"

type Pet struct {
	ID   uuid.UUID
	Name string
}

type DTO struct {
	Name string
}

func (d DTO) ToPet() Pet {
	return Pet{
		Name: d.Name,
	}
}
