package models

type Avaliacao struct {
	UserId      int
	Estrelas    int
	Comentarios []Comentario
}

type Comentario struct {
	Titulo   string
	Conteudo string
}
