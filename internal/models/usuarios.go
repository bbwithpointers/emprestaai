package models

type UserDefaultData struct {
	Nome                string      `json:"nome"`
	Localizacao         string      `json:"localizacao"`
	Endereco            string      `json:"endereco"`
	CEP                 int         `json:"cep"`
	Interesses          string      `json:"interesses"`
	Documento           string      `json:"documento"`
	Tipo                UsuarioTipo `json:"tipo"`
	NumeroDeEmprestimos int         `json:"numeroDeEmprestimos,omitempty"`
	Avaliacao           int         `json:"avaliacao,omitempty"`
}

type Usuario struct {
	UserDefaultData
	ID           int    `json:"id,omitempty"`
	DocumentoCPF string `json:"cpf"`
}
