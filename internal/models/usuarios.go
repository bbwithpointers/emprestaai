package models

const (
	CONTRATANTE int = 1
	TRABALHADOR int = 2
)

type Usuario struct {
	Nome                string
	Localizacao         Localizacao
	Endereco            string
	CEP                 int
	Interesses          []string
	Tipo                int
	NumeroDeEmprestimos int
	Avaliacao           Avaliacao
}

type Contratante struct {
	Usuario
	CPF string
}

type Trabalhador struct {
	Usuario
	CNPJ        string
	Ferramentas []Ferramentas
	// datetime
	Disponivel string
	// datetime
	ProximoHorario string
	ValorHora      int
}
