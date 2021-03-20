package commands

import (
	"jar-service/delivery/models/jarmodel"
)

type AddJarCmd struct {
	Jar *jarmodel.Jar `json:"jar" validate:"required"`
}
