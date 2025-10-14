package gen

import "github.com/google/uuid"

type UUIDGenerator func() uuid.UUID

func UUID() UUIDGenerator {
	return func() uuid.UUID {
		return uuid.New()
	}
}
