package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SignIn struct {
}

func NewSignIng() *SignIn {
	return &SignIn{}
}

func (s SignIn) Register(e *echo.Echo) {
	e.GET("/", Registro)
}

func Registro(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world\n")
}
