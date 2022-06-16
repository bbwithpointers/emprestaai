package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func StartApp() {

	e := echo.New()
	s := http.Server{
		Addr:    ":8080",
		Handler: e,
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Info(err)
	}
}
