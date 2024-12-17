package pet

import "context"

func (s service) Create(ctx context.Context, dto DTO) (*Pet, error) {
	if len(dto.Name) < 3 {
		return nil, ErrNameMinLength
	}

	return s.repository.Create(ctx, dto.ToPet())
}