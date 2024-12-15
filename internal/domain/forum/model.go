package forum

import "github.com/google/uuid"

type Forum struct {
	ID   uuid.UUID
	Name string
}

type DTO struct {
	Name string
}

func (d DTO) ToForum() Forum {
	return Forum{
		Name: d.Name,
	}
}
