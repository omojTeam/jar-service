package app

import "jar-service/domain"

type App struct {
	JarService domain.JarService
}

func NewApp(es domain.JarService) *App {
	return &App{
		JarService: es,
	}
}
