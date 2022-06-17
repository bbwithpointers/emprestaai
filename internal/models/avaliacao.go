package models

type Avaliacao struct {
	UserId      int          `json:"userId"`
	Estrelas    int          `json:"estrelas"`
	Comentarios []Comentario `json:"comentarios"`
}

type Comentario struct {
	Titulo   string `json:"titulo"`
	Conteudo string `json:"descricao"`
}
