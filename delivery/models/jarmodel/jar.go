package jarmodel

import "github.com/google/uuid"

type Jar struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title" validate:"required"`
	JarCode     string    `json:"jarCode"`
	TimesPerDay uint      `json:"timesPerDay" validate:"required"`
	Cards       []Card    `json:"cards" validate:"required"`
}
