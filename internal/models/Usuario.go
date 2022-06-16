package models

const (
	USUARIO     int = 1
	TRABALHADOR int = 2
)

type Usuario struct {
	Nome        string
	Localizacao int
	Endereco    string
	CEP         int
	Interesses  []string
	Tipo        int
}
