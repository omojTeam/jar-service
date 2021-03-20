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

	err := r.db.Create(&entry).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAllByJarCode(jarCode *string) (*domainmodel.Jar, error) {
	var entry domainmodel.Jar
	err := r.db.Preload("Cards").Where("jar_code = ?", jarCode).First(&entry).Error

	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func NewJarRepository(dbConn *gorm.DB) domain.JarRepository {

	return &repository{
		db: dbConn,
	}
}
