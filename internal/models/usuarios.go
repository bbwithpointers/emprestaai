package models

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
