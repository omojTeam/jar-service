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

func (js *jarService) GetJar(jarCode *string) (*responses.JarModel, error) {

	return nil, nil
}

func NewJarService(er domain.JarRepository) domain.JarService {
	es := &jarService{
		JarRepository: er,
	}

	return es
}
