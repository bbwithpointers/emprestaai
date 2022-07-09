package models

type UserDefaultData struct {
	Nome                string      `json:"nome"`
	Localizacao         Localizacao `json:"localizacao"`
	Endereco            string      `json:"endereco"`
	CEP                 int         `json:"cep"`
	Interesses          []string    `json:"interesses"`
	Documento           string      `json:"documento"`
	Tipo                UsuarioTipo `json:"tipo"`
	NumeroDeEmprestimos int         `json:"numeroDeEmprestimos,omitempty"`
	Avaliacao           Avaliacao   `json:"avaliacao,omitempty"`
}

type Usuario struct {
	UserDefaultData
	ID           string `json:"id,omitempty"`
	DocumentoCPF string `json:"cpf"`
}
