package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/matveevfg/AI-HR/backend/api/requests"
)

func (s *Server) Vacancy(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "malformatted id")
	}

	vacancy, err := s.service.Vacancy(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, vacancy)
}

func (s *Server) Vacancies(c echo.Context) error {
	var filter requests.VacancyFilter
	if err := c.Bind(&filter); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON body")
	}

	vacancies, err := s.service.Vacancies(c.Request().Context(), filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, vacancies)
}

func (s *Server) SaveVacancy(c echo.Context) error {
	var vacancyRequest requests.Vacancy
	if err := c.Bind(&vacancyRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON body")
	}

	id, err := s.service.SaveVacancy(c.Request().Context(), vacancyRequest.ToModel())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, id)
}

func (s *Server) UpdateVacancy(c echo.Context) error {
	var vacancyRequest requests.Vacancy
	if err := c.Bind(&vacancyRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid JSON body")
	}

	id, err := s.service.SaveVacancy(c.Request().Context(), vacancyRequest.ToModel())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, id)
}

func (s *Server) DeleteVacancy(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "malformatted id")
	}

	if err := s.service.DeleteVacancy(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (s *Server) SetVacancyInactive(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "malformatted id")
	}

	if err := s.service.SetVacancyInactive(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (s *Server) SetVacancyActive(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "malformatted id")
	}

	if err := s.service.SetVacancyActive(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
