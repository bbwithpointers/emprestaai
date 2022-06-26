package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/brunogbarros/emprestaai.git/internal/models"
	"github.com/brunogbarros/emprestaai.git/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
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
// var listaDeContratante []models.Contratante
var listaDeTrabalhador []models.Trabalhador

func Cadastro(c echo.Context) error {
	collections := repository.NewDBClient().Database("emprestaai").Collection("usuarios")
	ctx := context.Background()
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
			DocumentoCPF: u.Documento,
		}
		contra := []interface{}{
			bson.D{{"nome", contratante.Usuario.Nome}},
		}
		res, insertErr := collections.InsertMany(ctx, contra)
		if insertErr != nil {
			log.Errorf("insert error: %v", insertErr)
		}
		fmt.Println("response from insert: ", res)
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
	collection := repository.NewDBClient().Database("emprestaai").Collection("usuarios")
	ctx := context.Background()
	// vem do mongodb o listaAll e o ListaById
	// list vai no banco e devolve a lista
	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var contratantes []models.Contratante
	if err := cur.All(ctx, &contratantes); err != nil {
		panic(err)
	}
	fmt.Println(contratantes)
	return c.JSON(http.StatusOK, contratantes)
}

// func Avaliar() {
// 	// ao clicar em avaliar, o ID virá do usuário clicado
// 	idClicado :=
// }
