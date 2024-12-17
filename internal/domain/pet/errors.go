package pet

import "errors"

var (
	ErrNameMinLength = errors.New("name must be at least 3 characters")
)
