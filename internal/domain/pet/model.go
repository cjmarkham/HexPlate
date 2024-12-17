package pet

import "github.com/google/uuid"

type Pet struct {
	ID   uuid.UUID
	Name string
}

type DTO struct {
	ID   *uuid.UUID
	Name string
}

func (d DTO) ToPet() Pet {
	p := Pet{
		Name: d.Name,
	}

	if d.ID != nil {
		p.ID = *d.ID
	}

	return p
}
