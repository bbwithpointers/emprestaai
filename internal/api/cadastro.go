package api

import (
	"net/http"

	"github.com/brunogbarros/emprestaai.git/internal/models"
	"github.com/labstack/echo/v4"
)

type SignIn struct {
}

func NovoCadastro() *SignIn {
	return &SignIn{}
}

func (s SignIn) Register(e *echo.Echo) {
	e.POST("/cadastro", Cadastro)
}

func Cadastro(c echo.Context) error {

	u := new(models.Usuario)
	if err := c.Bind(u); err != nil {
		return err
	}
	var contratante models.Contratante
	var user models.Usuario
	if u.Tipo == 1 {
		user = models.Usuario{
			Nome:                u.Nome,
			Localizacao:         u.Localizacao,
			Endereco:            u.Endereco,
			CEP:                 u.CEP,
			Interesses:          u.Interesses,
			Documento:           u.Documento,
			Tipo:                u.Tipo,
			NumeroDeEmprestimos: u.NumeroDeEmprestimos,
			Avaliacao:           u.Avaliacao,
		}
		contratante = models.Contratante{
			Usuario:      user,
			DocumentoCPF: u.Documento,
		}
	}

	return c.JSON(http.StatusOK, contratante)
}
