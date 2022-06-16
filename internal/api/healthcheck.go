package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Healthcheck struct{}

func NewHealthcheck() *Healthcheck {
	return &Healthcheck{}
}

func (h *Healthcheck) Register(e *echo.Echo) {
	e.GET("/liveness", h.liveness)
	e.GET("/readness", h.readness)
}

func (h *Healthcheck) liveness(c echo.Context) error {
	r := make(map[string]string)
	r["status"] = "OK"
	return c.JSON(http.StatusOK, r)
}

func (h *Healthcheck) readness(c echo.Context) error {
	r := make(map[string]string)
	r["status"] = "OK"
	return c.JSON(http.StatusOK, r)
}
