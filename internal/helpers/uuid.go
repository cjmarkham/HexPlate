package helpers

//go:generate mockgen -source=uuid.go -package=helpers -destination=uuid_mock.go

import "github.com/google/uuid"

type UUIDGenerator interface {
	NewUUID() string
}

type DefaultUUIDGenerator struct{}

func (d *DefaultUUIDGenerator) NewUUID() *uuid.UUID {
	u := uuid.New()
	return &u
}

func ProvideUUIDGenerator() DefaultUUIDGenerator {
	return DefaultUUIDGenerator{}
}
