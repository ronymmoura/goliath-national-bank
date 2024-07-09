package token

import (
	"time"

	"github.com/google/uuid"
)

type Maker interface {
	CreateToken(id uuid.UUID, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
