package api

import (
	"fmt"
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
	user := models.Usuario{
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
	var contratante models.Contratante
	if u.Tipo == 1 && len(u.Documento) <= 11 {
		contratante = models.Contratante{
			Usuario:      user,
			DocumentoCPF: u.Documento,
		}
		// logica de salvar no banco
	}
	trabalhador := models.Trabalhador{
		Usuario:       user,
		DocumentoCNPJ: u.Documento,
		// default
		Disponivel: true,
	}
	fmt.Println(trabalhador)

	return c.JSON(http.StatusOK, contratante)
}
