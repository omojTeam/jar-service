package cron

import (
	"jar-service/app"
	"log"

	"github.com/robfig/cron"
)

type server struct {
	*app.App
}

func NewCron(app *app.App) {
	handler := &server{
		app,
	}

	handler.EnableCron()
}

func (s *server) EnableCron() {
	c := cron.New()
	c.AddFunc("@midnight", func() {
		log.Println("Resetting CardsSeenThisDay at midnight!")
		err := s.JarService.ResetCardsSeenThisDay()
		if err != nil {
			log.Println(err)
		} else {
			log.Print("Cards resetted succesfully!")
		}
	})
	c.Start()
}
