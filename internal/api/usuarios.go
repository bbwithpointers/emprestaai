package api

import (
	"github.com/brunogbarros/emprestaai.git/internal/repository"
	"github.com/brunogbarros/emprestaai.git/public"
	"html/template"
	"net/http"

	"github.com/brunogbarros/emprestaai.git/internal/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type Cadastro struct {
}

func NewCadastro() *Cadastro {
	return &Cadastro{}
}

func (s Cadastro) Register(e *echo.Echo) {
	t := &public.Template{
		Templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t
	// TODO: teste only
	e.GET("/cadastro", CriarCadastro)
	//e.GET("/listaTrabalhadores", ListarTrabalhadores)
	//e.GET("/listaContratante", ListarContratantes)
	//e.GET("/hello", Hello)
}
func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "world")
}

// temporario
var listaDeLoja []models.Loja

func CriarCadastro(c echo.Context) error {

	u := new(models.Usuario)
	u.Nome = "bruno"
	u.DocumentoCPF = "04038792004"
	u.Localizacao = "123123,1312312"
	u.Endereco = "rua tal casa"
	u.CEP = 98787665
	u.Interesses = "go, dormir"
	u.Tipo = models.USUARIO
	u.Avaliacao = 123
	//if err := c.Bind(u); err != nil {
	//	return err
	//}
	user := models.Usuario{
		Nome:                u.Nome,
		Localizacao:         u.Localizacao,
		DocumentoCPF:        u.DocumentoCPF,
		Endereco:            u.Endereco,
		CEP:                 u.CEP,
		Interesses:          u.Interesses,
		Tipo:                u.Tipo,
		NumeroDeEmprestimos: 0,
		Avaliacao:           u.Avaliacao,
	}
	//var contratante models.Usuario
	//if u.Tipo == models.USUARIO && len(u.Documento) <= 11 {
	//	contratante = models.Usuario{
	//		Usuario: user,
	//		DocumentoCPF:    u.Documento,
	//	}
	//	fmt.Println("CONTRATANTE: ", contratante)
	repository.NovoRepositorioDeUsuario()

	//} else {
	//	loja := models.Loja{
	//		ID:            "223-223",
	//		DocumentoCNPJ: u.Documento,
	//		// default
	//		Disponivel: true,
	//	}
	//	listaDeLoja = append(listaDeLoja, loja)
	//	fmt.Println(loja)
	//	return c.JSON(http.StatusOK, user)
	//}

	return c.JSON(http.StatusOK, user)
}

func ListarTrabalhadores(c echo.Context) error {
	// vem do mongodb o listaAll e o ListaById
	// list vai no banco e devolve a lista
	if len(listaDeLoja) == 0 {
		return c.JSON(http.StatusOK, "Lista vazia")
	}

	return c.JSON(http.StatusOK, listaDeLoja)
}

func ListarContratantes(c echo.Context) error {
	return c.JSON(http.StatusOK, "contratantes")
}
