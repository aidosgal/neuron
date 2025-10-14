package gen

import (
	"fmt"

	"github.com/google/uuid"
)

type UUIDGenerator func() uuid.UUID

func UUID() UUIDGenerator {
	return func() uuid.UUID {
		return uuid.Must(uuid.New(), fmt.Errorf("failed to generate uuid"))
	}
}
