package domain

import (
	"jar-service/delivery/commands"
	"jar-service/delivery/responses"
	"jar-service/domain/domainmodel"
)

type JarService interface {
	AddJar(AddJarCmd *commands.AddJarCmd) (*string, error)
	GetJar(jarCode *string) (*responses.JarModel, error)
}

type JarRepository interface {
	Create(jar *domainmodel.Jar) error
	GetByJarCode(jarCode *string) (*domainmodel.Jar, error)
}
