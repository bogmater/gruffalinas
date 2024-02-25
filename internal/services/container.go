package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Container struct {
	Config    *Config
	Web       *echo.Echo
	Validator *Validator
	Db        *pgxpool.Pool
}

func NewContainer() *Container {
	c := new(Container)

	c.initConfig()
	c.initWeb()
	c.initValidator()
	c.initDatabase()

	return c
}

func (c *Container) initConfig() {
	cfg, err := GetConfig()

	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	c.Config = &cfg
}

func (c *Container) initWeb() {
	c.Web = echo.New()

	c.Web.Static("/", "public")

	switch c.Config.App.Environment {
	case EnvProduction:
		c.Web.Logger.SetLevel(log.WARN)
	default:
		c.Web.Logger.SetLevel(log.DEBUG)
	}
}

func (c *Container) initValidator() {
	c.Validator = NewValidator()
}

func (c *Container) initDatabase() {
	var err error

	getAddr := func(dbName string) string {
		return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
			c.Config.Database.User,
			c.Config.Database.Password,
			c.Config.Database.Hostname,
			c.Config.Database.Port,
			dbName,
		)
	}

	c.Db, err = pgxpool.New(context.Background(), getAddr(c.Config.Database.Database))

	if err != nil {
		panic(fmt.Sprintf("failed to connect to database %v", err))
	}
}

func (c *Container) Shutdown() error {
	c.Db.Close()

	return nil
}
