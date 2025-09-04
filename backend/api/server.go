package api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	e *echo.Echo

	service AiHrService
}

func New(service AiHrService) *Server {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	s := &Server{
		e:       e,
		service: service,
	}

	s.setupRoutes()

	return s
}

func (s *Server) setupRoutes() {
	api := s.e.Group("/api")

	api.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	})

	vacancies := api.Group("/vacancies")
	vacancies.GET("", s.Vacancies)
	vacancies.GET("/:id", s.Vacancy)
	vacancies.POST("", s.SaveVacancy)
	vacancies.PUT("/:id", s.UpdateVacancy)
	vacancies.DELETE("/:id", s.DeleteVacancy)
	vacancies.PUT("/:id/active", s.SetVacancyActive)
	vacancies.PUT("/:id/inactive", s.SetVacancyInactive)

	resumes := api.Group("/resumes")
	resumes.POST("/:vacancy-id", s.UploadResumes)
	resumes.GET("/:vacancy-id", s.Resumes)

	dialogs := api.Group("/dialogs")
	dialogs.GET("/ws", s.handleWebSocket)
}

func (s *Server) Start(addr string) error {
	return s.e.Start(addr)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}
