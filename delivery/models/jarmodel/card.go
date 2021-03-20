package jarmodel

import "github.com/google/uuid"

type Card struct {
	ID   uuid.UUID `json:"id"`
	Text string    `json:"text" validate:"required"`
}
