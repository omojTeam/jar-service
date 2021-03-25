package domainmodel

import (
	"jar-service/delivery/models/jarmodel"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Card struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `sql:"index"`
	Text      string
	Seen      bool
	JarID     uuid.UUID
}

func (Card) TableName() string {
	return "jar_service.cards"
}

func (question *Card) BeforeCreate(tx *gorm.DB) (err error) {
	nullUUID := uuid.UUID{}
	if question.ID == nullUUID {
		question.ID = uuid.New()
	}
	return
}

func newCard(c *jarmodel.Card, jID uuid.UUID) *Card {
	return &Card{
		ID:    c.ID,
		Text:  c.Text,
		JarID: jID,
	}
}

func newCardArray(cArr *[]jarmodel.Card, jID uuid.UUID) *[]Card {
	cards := []Card{}
	for _, q := range *cArr {
		cards = append(cards, *newCard(&q, jID))
	}

	return &cards
}
