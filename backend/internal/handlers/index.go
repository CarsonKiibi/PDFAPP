package handlers

import (
	"github.com/carsonkiibi/PDFAPP/backend/internal/templates"
	"github.com/labstack/echo/v4"
)

func HandleIndex(c echo.Context) error {
	return templates.Index().Render(c.Request().Context(), c.Response().Writer)
}
