package repository

import (
	"database/sql"
	"fmt"
	"github.com/brunogbarros/emprestaai.git/internal/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// TODO: persistir
func ConnectDB() {
	db := dbConn()
	// important for mysql
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "meemprestaai"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Insert(usuario models.Usuario) {
	db := dbConn()
	stmt, err := db.Prepare("INSERT INTO tb_usuarios(null, nome,localizacao,cep,interesses,documento,tipo,numeroDeEmprestimos,avaliacaoId, endereco) VALUES (?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	result, _ := stmt.Exec(usuario.Nome, usuario.Localizacao, usuario.CEP, usuario.Interesses, usuario.Documento, usuario.Tipo, usuario.NumeroDeEmprestimos, usuario.Avaliacao.UserId, usuario.Endereco)
	lastIdInserted, _ := result.LastInsertId()
	fmt.Printf("INSERTED %s", usuario.Nome, " id: %d", int(lastIdInserted))
	defer db.Close()
}
