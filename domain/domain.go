package domain

import (
	"jar-service/delivery/commands"
	"jar-service/delivery/responses"
	"jar-service/domain/domainmodel"
)

type JarService interface {
	AddJar(AddJarCmd *commands.AddJarCmd) (*string, error)
	GetAllJar(jarCode *string) (*responses.JarModel, error)
	GetOneCard(jarCode *string) (*responses.JarModel, error)
	ResetJar(jarCode *string) error
	ResetCardsSeenThisDay() error
}

type JarRepository interface {
	Create(jar *domainmodel.Jar) error
	GetAllByJarCode(jarCode *string) (*domainmodel.Jar, error)
	GetOneCardByJarCode(jarCode *string) (*domainmodel.Jar, error)
	UpdateJar(jar *domainmodel.Jar) error
	ResetCardsSeenThisDay() error
}
