package http

import (
	"jar-service/app"
	"jar-service/delivery/commands"
	"jar-service/domain"
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
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	if err = c.Validate(cmd); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
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

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Error(err)

	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound, domain.ErrRecordNotFound, gorm.ErrRecordNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	case domain.ErrUnauthorized:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
