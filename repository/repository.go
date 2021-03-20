package repository

import (
	"jar-service/domain"
	"jar-service/domain/domainmodel"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(entry *domainmodel.Jar) error {

	return nil
}

func (r *repository) GetByJarCode(jarCode *string) (*domainmodel.Jar, error) {

	return nil, nil
}

func NewJarRepository(dbConn *gorm.DB) domain.JarRepository {

	return &repository{
		db: dbConn,
	}
}
