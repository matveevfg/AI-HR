package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *Server) UploadResumes(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	files := form.File["resume"]

	vacancyID, err := uuid.Parse(c.Param("vacancy-id"))

	if err := s.service.SaveResume(c.Request().Context(), files, vacancyID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}

func (s *Server) Resumes(c echo.Context) error {
	vacancyID, err := uuid.Parse(c.Param("vacancy-id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resumes, err := s.service.Resumes(c.Request().Context(), vacancyID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resumes)
}
