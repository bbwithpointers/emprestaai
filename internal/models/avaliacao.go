package models

type Avaliacao struct {
	UserId      string       `json:"userId"`
	Estrelas    int          `json:"estrelas"`
	Comentarios []Comentario `json:"comentarios"`
}

type Comentario struct {
	Titulo   string `json:"titulo,omitempty"`
	Conteudo string `json:"descricao,omitempty"`
}
