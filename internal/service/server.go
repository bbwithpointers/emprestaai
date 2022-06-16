package service

import (
	"github.com/brunogbarros/emprestaai.git/internal/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func NewApp() *echo.Echo {

	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.RemoveTrailingSlash())
	app.Use(middleware.CORS())

	api.NewSignIng().Register(app)
	return app

}
func NewHealthApp() *echo.Echo {
	app := echo.New()

	api.NewHealthcheck().Register(app)
	return app
}

func Start(e *echo.Echo, host string) {

	err := e.Start(host)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Server start at port %s", host)
}
