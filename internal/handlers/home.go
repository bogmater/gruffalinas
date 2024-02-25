package handlers

import (
	"context"
	"gruffalinas/internal/services"
	"gruffalinas/templates"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c *services.Container) echo.HandlerFunc {
	return func(e echo.Context) error {

		// var result string
		//
		// result, err := c.db.selecthellodatabase(context.background())
		//
		// if err != nil {
		// 	e.logger().fatal(err)
		// }
		//
		// content := layout.main()

		// return templates.layout(content).render(context.background(), e.response().writer)

		return templates.Layout("Home", templates.Home()).
			Render(context.Background(), e.Response().Writer)
	}
}
