package jarmodel

import (
	"github.com/google/uuid"
)

type Jar struct {
	ID             uuid.UUID `json:"id"`
	Title          string    `json:"title" validate:"required"`
	CardsPerDay    uint      `json:"cardsPerDay" validate:"required"`
	CardsLeftToday uint      `json:"cardsLeftToday"`
	CardsLeft      uint      `json:"cardsLeft"`
	RecipientEmail string    `json:"recipientEmail" validate:"required,email"`
	Cards          []Card    `json:"cards" validate:"required"`
}
