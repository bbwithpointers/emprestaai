package models

const (
	CONTRATANTE int = 1
	TRABALHADOR int = 2
)

type Usuario struct {
	Nome                string      `json:"nome"`
	Localizacao         Localizacao `json:"localizacao"`
	Endereco            string      `json:"endereco"`
	CEP                 int         `json:"cep"`
	Interesses          []string    `json:"interesses"`
	Documento           string      `json:"documento"`
	Tipo                int         `json:"tipo"`
	NumeroDeEmprestimos int         `json:"numeroDeEmprestimos,omitempty"`
	Avaliacao           Avaliacao   `json:"avaliacao,omitempty"`
}

type Contratante struct {
	Usuario
	ID           string `json:"id,omitempty"`
	DocumentoCPF string `json:"cpf"`
}

// mover para um struct proprio
type Loja struct {
	ID            string        `json:"id,omitempty"`
	DocumentoCNPJ string        `json:"cnpj"`
	Ferramentas   []Ferramentas `json:"ferramentas,omitempty"`
	// datetime
	Disponivel bool `json:"disponivel"`
	// datetime
	ProximoHorario string    `json:"proximoHorario,omitempty"`
	ValorHora      int       `json:"valorHora,omitempty"`
	Avaliacao      Avaliacao `json:"avaliacao,omitempty"`
}
