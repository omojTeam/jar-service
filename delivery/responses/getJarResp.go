package responses

import (
	"jar-service/delivery/models/jarmodel"
	"jar-service/domain/domainmodel"

	"github.com/google/uuid"
)

type JarModel struct {
	Jar *jarmodel.Jar `json:"jar"`
}

func NewJarModelResp(jar *domainmodel.Jar) *JarModel {
	return &JarModel{
		Jar: newJarModel(jar),
	}
}

func newJarModel(domainJar *domainmodel.Jar) *jarmodel.Jar {

	return &jarmodel.Jar{
		ID:             domainJar.ID,
		Title:          domainJar.Title,
		CardsPerDay:    domainJar.CardsPerDay,
		CardsLeftToday: domainJar.CardsPerDay - domainJar.CardsSeenThisDay,
		CardsLeft:      domainJar.NumOfCards - domainJar.CardsSeen,
		RecipientEmail: domainJar.RecipientEmail,
		Cards:          *newJarCardArray(&domainJar.Cards, domainJar.ID),
	}
}

func newJarCard(c *domainmodel.Card, tID uuid.UUID) *jarmodel.Card {
	return &jarmodel.Card{
		ID:   c.ID,
		Text: c.Text,
	}
}

func newJarCardArray(cArr *[]domainmodel.Card, jID uuid.UUID) *[]jarmodel.Card {
	cards := []jarmodel.Card{}
	for _, q := range *cArr {
		cards = append(cards, *newJarCard(&q, jID))
	}

	return &cards
}
