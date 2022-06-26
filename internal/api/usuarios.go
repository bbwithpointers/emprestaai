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
	e.GET("/listaTrabalhadores", ListarTrabalhadores)
	e.GET("/listaContratante", ListarContratantes)
}

// temporario
var listaDeContratante []models.Contratante
var listaDeTrabalhador []models.Trabalhador

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
		NumeroDeEmprestimos: 0,
		Avaliacao:           models.Avaliacao{},
	}
	var contratante models.Contratante
	if u.Tipo == 1 && len(u.Documento) <= 11 {
		contratante = models.Contratante{
			Usuario:      user,
			ID:           "123-123",
			DocumentoCPF: u.Documento,
		}
		listaDeContratante = append(listaDeContratante, contratante)
	} else {
		trabalhador := models.Trabalhador{
			Usuario:       user,
			ID:            "223-223",
			DocumentoCNPJ: u.Documento,
			// default
			Disponivel: true,
		}
		listaDeTrabalhador = append(listaDeTrabalhador, trabalhador)
		fmt.Println(trabalhador)
		return c.JSON(http.StatusOK, trabalhador)
	}

	return c.JSON(http.StatusOK, contratante)
}

func ListarTrabalhadores(c echo.Context) error {
	// vem do mongodb o listaAll e o ListaById
	// list vai no banco e devolve a lista
	if len(listaDeTrabalhador) == 0 {
		return c.JSON(http.StatusOK, "Lista vazia")
	}

	return c.JSON(http.StatusOK, listaDeTrabalhador)
}

func ListarContratantes(c echo.Context) error {
	// vem do mongodb o listaAll e o ListaById
	// list vai no banco e devolve a lista
	return c.JSON(http.StatusOK, listaDeContratante)
}

// func Avaliar() {
// 	// ao clicar em avaliar, o ID virá do usuário clicado
// 	idClicado :=
// }
