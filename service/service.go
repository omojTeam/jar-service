package service

import (
	"jar-service/delivery/commands"
	"jar-service/delivery/responses"
	"jar-service/domain"
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

	if result.CardsSeen >= result.NumOfCards || int(result.CardsPerDay)-int(result.CardsSeenThisDay) <= 0 {
		return nil, domain.ErrNotFound
	}

	if len(result.Cards) <= 0 {
		return nil, domain.ErrRecordNotFound
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

func NewJarService(er domain.JarRepository) domain.JarService {
	es := &jarService{
		JarRepository: er,
	}

	return es
}
