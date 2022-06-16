package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SignIn struct {
}

func NovoCadastro() *SignIn {
	return &SignIn{}
}

func (s SignIn) Register(e *echo.Echo) {
	e.GET("/cadastro", Cadastro)
}

func Cadastro(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world\n")
}
