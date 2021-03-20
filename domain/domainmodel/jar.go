package domainmodel

import (
	"errors"
	"jar-service/delivery/commands"
	"jar-service/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Jar struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `sql:"index"`
	Title          string
	JarCode        string `qorm:"unique"`
	TimesPerDay    uint
	NumOfCards     uint
	RecipientEmail string
	Cards          []Card `gorm:"foreignKey:JarID"`
}

func (Jar) TableName() string {
	return "jar_service.jar"
}

func (jar *Jar) BeforeCreate(tx *gorm.DB) (err error) {
	nullUUID := uuid.UUID{}
	if jar.ID == nullUUID {
		jar.ID = uuid.New()
	}
	return
}

func NewJarModel(cmd *commands.AddJarCmd) (*Jar, error) {

	if err := validateCommand(cmd); err != nil {
		return nil, err
	}

	return &Jar{
		Title:          cmd.Jar.Title,
		JarCode:        utils.RandomCode(10),
		TimesPerDay:    cmd.Jar.TimesPerDay,
		RecipientEmail: cmd.Jar.RecipientEmail,
		NumOfCards:     uint(len(cmd.Jar.Cards)),
		Cards:          *newCardArray(&cmd.Jar.Cards, cmd.Jar.ID),
	}, nil
}

func validateCommand(cmd *commands.AddJarCmd) error {
	if &cmd.Jar.Title == nil {
		return errors.New("Title can not be empty!")
	}

	if &cmd.Jar.TimesPerDay == nil {
		return errors.New("TimesPerDay can not be empty!")
	}

	if &cmd.Jar.RecipientEmail == nil {
		return errors.New("RecipientEmail can not be empty!")
	}

	if &cmd.Jar.Cards == nil {
		return errors.New("Questions can not be empty!")
	}

	return nil
}
