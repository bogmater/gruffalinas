package routes

import (
	"gruffalinas/internal/handlers"
	"gruffalinas/internal/middleware"
	"gruffalinas/internal/services"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

func BuildRoutes(c *services.Container) {
	g := c.Web.Group("")

	g.Use(
		echomw.RemoveTrailingSlashWithConfig(echomw.TrailingSlashConfig{
			RedirectCode: http.StatusMovedPermanently,
		}),
		echomw.Recover(),
		echomw.Secure(),
		echomw.RequestID(),
		echomw.Gzip(),
		echomw.Logger(),
		middleware.LogRequestID(),
		session.Middleware(sessions.NewCookieStore([]byte(c.Config.App.EncryptionKey))),
		echomw.CSRFWithConfig(echomw.CSRFConfig{
			TokenLookup: "form:csrf",
		}),
	)

	AppRoutes(c, g)
}

func AppRoutes(c *services.Container, g *echo.Group) {
	g.GET("/", handlers.HomeHandler(c))
	g.GET("/login", handlers.LoginHandler(c))
}
