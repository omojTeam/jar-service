package http

import (
	"jar-service/app"
	"jar-service/delivery/commands"
	"jar-service/domain/domainerrors"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

type ResponseJarCode struct {
	JarCode string `json:"jarCode"`
}

type server struct {
	*app.App
}

func NewHandler(e *echo.Echo, app *app.App) {
	handler := &server{
		app,
	}
	e.GET("/", handler.Index)
	e.POST("/jar", handler.AddJar)
	e.GET("/jar/card/:code", handler.GetCard)
	e.GET("/jar/reset/:code", handler.ResetJar)

	//TODO: Temp for debugging purposes
	e.GET("/jar/all/:code", handler.GetJar)
	e.GET("/jar/resetall", handler.ResetCardsSeenThisDay)
}

func (s *server) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseMessage{Message: "OK"})
}

func (s *server) AddJar(c echo.Context) error {
	var cmd commands.AddJarCmd

	err := c.Bind(&cmd)
	if err != nil {
		return c.JSON(getStatusCode(domainerrors.ErrUnprocessableEntity), ResponseMessage{Message: err.Error()})
	}

	if err = c.Validate(cmd); err != nil {
		return c.JSON(getStatusCode(domainerrors.ErrUnprocessableEntity), ResponseMessage{Message: err.Error()})
	}

	jarCode, err := s.JarService.AddJar(&cmd)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseJarCode{JarCode: *jarCode})
}

func (s *server) GetJar(c echo.Context) error {
	var jarCode = c.Param("code")

	resp, err := s.JarService.GetAllJar(&jarCode)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *server) GetCard(c echo.Context) error {
	var jarCode = c.Param("code")

	resp, err := s.JarService.GetOneCard(&jarCode)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *server) ResetCardsSeenThisDay(c echo.Context) error {

	err := s.JarService.ResetCardsSeenThisDay()

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseMessage{Message: "ok"})
}

func (s *server) ResetJar(c echo.Context) error {
	var jarCode = c.Param("code")

	err := s.JarService.ResetJar(&jarCode)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseMessage{Message: "ok"})
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Error(err)

	switch err {
	case domainerrors.ErrInternalServerError:
		return http.StatusInternalServerError
	case domainerrors.ErrNotFound, domainerrors.ErrRecordNotFound, gorm.ErrRecordNotFound, domainerrors.ErrNoCardsLeft:
		return http.StatusNotFound
	case domainerrors.ErrConflict:
		return http.StatusConflict
	case domainerrors.ErrUnauthorized:
		return http.StatusUnauthorized
	case domainerrors.ErrBadParamInput, domainerrors.ErrUnprocessableEntity, domainerrors.ErrEmptyTitle, domainerrors.ErrEmptyCardsPerDay, domainerrors.ErrEmptyRecipientEmail, domainerrors.ErrEmptyCardArray, domainerrors.ErrCardsPerDayTooLarge:
		return http.StatusUnprocessableEntity
	case domainerrors.ErrForbidden, domainerrors.ErrNoCardsLeftToday:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
