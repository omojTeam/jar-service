package service

import (
	"jar-service/delivery/commands"
	"jar-service/delivery/responses"
	"jar-service/domain"
)

type jarService struct {
	JarRepository domain.JarRepository
}

func (js *jarService) AddJar(cmd *commands.AddJarCmd) (*string, error) {

	return nil, nil
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
