package jarmodel

import (
	"github.com/google/uuid"
)

type Jar struct {
	ID             uuid.UUID `json:"id"`
	Title          string    `json:"title" validate:"required"`
	TimesPerDay    uint      `json:"timesPerDay" validate:"required"`
	RecipientEmail string    `json:"recipientEmail" validate:"required,email"`
	Cards          []Card    `json:"cards" validate:"required"`
}
