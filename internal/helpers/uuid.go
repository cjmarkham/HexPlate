package helpers

//go:generate mockgen -source=uuid.go -package=helpers -destination=uuid_mock.go

import "github.com/google/uuid"

type UUIDGenerator interface {
	NewUUID() uuid.UUID
}

type Generator struct{}

func (d Generator) NewUUID() uuid.UUID {
	u := uuid.New()
	return u
}

func ProvideUUIDGenerator() UUIDGenerator {
	return Generator{}
}
