package domain

import "github.com/google/uuid"

type Resource struct {
	ID     uuid.UUID
	Format string
	Link   string
	UserID uuid.UUID
}
