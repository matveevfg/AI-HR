package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) UploadResume(c echo.Context) error {
	file, err := c.FormFile("resume")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "resume field is required")
	}

	if err := s.service.SaveResume(c.Request().Context(), file); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}
