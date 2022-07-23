package repository

import (
	"database/sql"
	"github.com/brunogbarros/emprestaai.git/internal/models"
)

type usuariosRepository struct {
	db *sql.DB
}

func NovoRepositorioDeUsuario(db *sql.DB) *usuariosRepository {
	return &usuariosRepository{db}
}

func (repo usuariosRepository) Criar(usuario models.Usuario) (uint64, error) {
	stmt, err := repo.db.Prepare("insert into tb_usuarios (nome,cpf,localizacao,endereco,cep,interesses,tipo,numeroDeEmprestimos, avaliacao) VALUES(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(usuario.Nome, usuario.DocumentoCPF, usuario.Localizacao, usuario.Endereco, usuario.CEP, usuario.Interesses, usuario.Tipo, usuario.NumeroDeEmprestimos, usuario.Avaliacao)
	if err != nil {
		return 0, err
	}
	lastIdInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastIdInserted), nil
}
