package handlers

import (
	"context"
	"gruffalinas/internal/services"
	"gruffalinas/templates"

	"github.com/labstack/echo/v4"
)

func LoginHandler(c *services.Container) echo.HandlerFunc {
	return func(e echo.Context) error {

		return templates.Layout("Login", templates.Login()).
			Render(context.Background(), e.Response().Writer)
	}
}
