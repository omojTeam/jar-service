package main

import (
	"context"
	"jar-service/app"
	"jar-service/config"
	"jar-service/delivery/http"
	"jar-service/domain/domainmodel"
	"jar-service/repository"
	"jar-service/service"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

var postgresDB *gorm.DB

func main() {

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{"*"},
		AllowCredentials: true,
	}))

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		logrus.Infof("Received %s signal", <-c)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	er := repository.NewJarRepository(postgresDB)
	es := service.NewJarService(er)
	a := app.NewApp(es)

	http.NewHandler(e, a)

	log.Fatal(e.Start(":" + config.Cfg.Port))
}

func init() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v\n", err)
	}

	postgresDB = initPostgres()
}

func initPostgres() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Cfg.DbAccess), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return migrate(db)
}

func migrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&domainmodel.Card{})
	db.AutoMigrate(&domainmodel.Jar{})

	return db
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
