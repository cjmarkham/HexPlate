package forum

import (
	"github.com/cjmarkham/hexplate/cmd/api"
	"github.com/cjmarkham/hexplate/internal/domain/forum"
)

func RequestToDTO(request api.CreateForumRequestObject) forum.DTO {
	return forum.DTO{
		Name: request.Body.Name,
	}
}
