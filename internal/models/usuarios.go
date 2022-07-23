package models

type Usuario struct {
	ID                  int         `json:"id,omitempty"`
	Nome                string      `json:"nome"`
	DocumentoCPF        string      `json:"cpf"`
	Localizacao         string      `json:"localizacao"`
	Endereco            string      `json:"endereco"`
	CEP                 int         `json:"cep"`
	Interesses          string      `json:"interesses"`
	Tipo                UsuarioTipo `json:"tipo"`
	NumeroDeEmprestimos int         `json:"numeroDeEmprestimos,omitempty"`
	Avaliacao           int         `json:"avaliacao,omitempty"`
}
