package domainmodel

import (
	"crypto/tls"
	"errors"
	"fmt"
	"jar-service/config"
	"jar-service/delivery/commands"
	"jar-service/utils"
	"time"

	"github.com/google/uuid"
	gomail "gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

type Jar struct {
	ID               uuid.UUID `gorm:"type:uuid;primary_key"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `sql:"index"`
	Title            string
	JarCode          string `qorm:"unique"`
	CardsPerDay      uint
	CardsSeen        uint
	CardsSeenThisDay uint
	NumOfCards       uint
	RecipientEmail   string
	Cards            []Card `gorm:"foreignKey:JarID"`
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
		CardsPerDay:    cmd.Jar.CardsPerDay,
		RecipientEmail: cmd.Jar.RecipientEmail,
		NumOfCards:     uint(len(cmd.Jar.Cards)),
		Cards:          *newCardArray(&cmd.Jar.Cards, cmd.Jar.ID),
	}, nil
}

func validateCommand(cmd *commands.AddJarCmd) error {
	if &cmd.Jar.Title == nil {
		return errors.New("Title can not be empty!")
	}

	if &cmd.Jar.CardsPerDay == nil {
		return errors.New("CardsPerDay can not be empty!")
	}

	if &cmd.Jar.RecipientEmail == nil {
		return errors.New("RecipientEmail can not be empty!")
	}

	if &cmd.Jar.Cards == nil {
		return errors.New("Questions can not be empty!")
	}

	return nil
}

func (jar *Jar) SendEmail() error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.Cfg.EmailUsername)
	m.SetHeader("To", jar.RecipientEmail)
	m.SetHeader("Subject", fmt.Sprintf("You've received a Happiness Jar! - %s", jar.Title))
	m.SetBody("text/html", fmt.Sprintf(`<html><body style="text-align:center;"><a href="%s%s"><img src="%s"></img></a></body></html>`, config.Cfg.GuiUrl, jar.JarCode, config.Cfg.EmailImageUrl))

	d := gomail.NewDialer(config.Cfg.EmailHost, config.Cfg.EmailPort, config.Cfg.EmailUsername, config.Cfg.EmailPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
