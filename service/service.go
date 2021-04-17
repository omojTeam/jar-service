package service

import (
	"jar-service/delivery/commands"
	"jar-service/delivery/responses"
	"jar-service/domain"
	"jar-service/domain/domainerrors"
	"jar-service/domain/domainmodel"
)

type jarService struct {
	JarRepository domain.JarRepository
}

func (js *jarService) AddJar(cmd *commands.AddJarCmd) (*string, error) {
	var err error

	jar, err := domainmodel.NewJarModel(cmd)
	if err != nil {
		return nil, err
	}

	err = js.JarRepository.Create(jar)
	if err != nil {
		return nil, err
	}

	err = jar.SendEmail()
	if err != nil {
		return nil, err
	}

	return &jar.JarCode, nil
}

func (js *jarService) GetAllJar(jarCode *string) (*responses.JarModel, error) {

	result, err := js.JarRepository.GetAllByJarCode(jarCode)
	if err != nil {
		return nil, err
	}

	return responses.NewJarModelResp(result), nil
}

func (js *jarService) GetOneCard(jarCode *string) (*responses.JarModel, error) {
	result, err := js.JarRepository.GetOneCardByJarCode(jarCode)
	if err != nil {
		return nil, err
	}

	if result.CardsSeen >= result.NumOfCards {
		return nil, domainerrors.ErrNoCardsLeft
	}

	if result.CardsPerDay-result.CardsSeenThisDay <= 0 {
		return nil, domainerrors.ErrNoCardsLeftToday
	}

	if len(result.Cards) <= 0 {
		return nil, domainerrors.ErrNoCardsLeft
	}

	result.Cards[0].Seen = true
	result.CardsSeenThisDay++
	result.CardsSeen++
	err = js.JarRepository.UpdateJar(result)
	if err != nil {
		return nil, err
	}

	return responses.NewJarModelResp(result), nil
}

func (js *jarService) ResetCardsSeenThisDay() error {
	err := js.JarRepository.ResetCardsSeenThisDay()
	if err != nil {
		return err
	}

	return nil
}

func (js *jarService) ResetJar(jarCode *string) error {

	result, err := js.JarRepository.GetAllByJarCode(jarCode)
	if err != nil {
		return err
	}

	if result.NumOfCards-result.CardsSeen <= 0 {

		result.CardsSeenThisDay = 0
		result.CardsSeen = 0
		for i := range result.Cards {
			result.Cards[i].Seen = false
		}

		err := js.JarRepository.UpdateJar(result)
		if err != nil {
			return err
		}
		return nil
	}

	return domainerrors.ErrBadParamInput
}

func NewJarService(er domain.JarRepository) domain.JarService {
	es := &jarService{
		JarRepository: er,
	}

	return es
}
