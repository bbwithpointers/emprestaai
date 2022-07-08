package models

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
