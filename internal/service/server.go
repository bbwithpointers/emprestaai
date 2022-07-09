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

	api.NewCadastro().Register(app)
	return app

}
func Start(e *echo.Echo, host string) {
	e.HideBanner = true
	err := e.Start(host)
	if err != nil {
		log.Fatal("erro ao inicializar o servidor: ", err)
	}
}
