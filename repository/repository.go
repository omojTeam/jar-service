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

func (r *repository) GetOneCardByJarCode(jarCode *string) (*domainmodel.Jar, error) {
	var entry domainmodel.Jar
	err := r.db.Preload("Cards",
		"seen = false",
		func(db *gorm.DB) *gorm.DB {
			return db.Limit(1)
		},
		func(db *gorm.DB) *gorm.DB {
			return db.Order("RANDOM()")
		}).Where("jar_code = ?", jarCode).Take(&entry).Error

	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (r *repository) UpdateJar(entry *domainmodel.Jar) error {

	for i := range entry.Cards {
		err := r.db.Model(&entry.Cards[i]).Update("Seen", entry.Cards[i].Seen).Error
		if err != nil {
			return err
		}
	}

	err := r.db.Model(&entry).Update("CardsSeen", entry.CardsSeen).Error
	if err != nil {
		return err
	}

	err = r.db.Model(&entry).Update("CardsSeenThisDay", entry.CardsSeenThisDay).Error
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

func (r *repository) ResetCardsSeenThisDay() error {
	err := r.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&domainmodel.Jar{}).Update("CardsSeenThisDay", 0).Error
	if err != nil {
		return err
	}

	return nil
}

func NewJarRepository(dbConn *gorm.DB) domain.JarRepository {

	return &repository{
		db: dbConn,
	}
}
